package usecase

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"

	"github.com/alexgtn/go-linkshort/internal/domain/nordpool"
)

type mockLinkRepo struct {
	maxID int
	links []*nordpool.DayAheadValue
}

func newMockRepo() *mockLinkRepo {
	return &mockLinkRepo{}
}

func newMockRepoWithSeed(links ...*nordpool.DayAheadValue) *mockLinkRepo {
	var lnks []*nordpool.DayAheadValue

	for _, l := range links {
		lnks = append(lnks, l)
	}

	return &mockLinkRepo{links: lnks}
}

func (m *mockLinkRepo) Create(ctx context.Context, long string) (*nordpool.DayAheadValue, error) {
	_, ok := lo.Find(m.links, func(l *nordpool.DayAheadValue) bool { return l.LongURL() == long })
	if ok {
		return nil, fmt.Errorf("link already exists %s", long)
	}

	m.maxID++
	l, _ := nordpool.NewLink(m.maxID, long, time.Now())
	m.links = append(m.links, l)

	return l, nil
}

func (m *mockLinkRepo) GetOneByLongURL(ctx context.Context, long string) (*nordpool.DayAheadValue, error) {
	l, ok := lo.Find(m.links, func(l *nordpool.DayAheadValue) bool { return l.LongURL() == long })
	if !ok {
		return nil, fmt.Errorf("link not found %s", long)
	}
	return l, nil
}

func (m *mockLinkRepo) GetOneByShortPath(ctx context.Context, short string) (*nordpool.DayAheadValue, error) {
	l, ok := lo.Find(m.links, func(l *nordpool.DayAheadValue) bool { return l.ShortPath() == short })
	if !ok {
		return nil, fmt.Errorf("link not found %s", short)
	}
	return l, nil
}

func (m *mockLinkRepo) SetShortPath(ctx context.Context, id int, path string) (*nordpool.DayAheadValue, error) {
	l, ok := lo.Find(m.links, func(l *nordpool.DayAheadValue) bool { return l.ID() == id })
	if !ok {
		return nil, fmt.Errorf("link not found %s", long)
	}

	l.SetShortPath(path)

	return l, nil
}

type linkService interface {
	Redirect(ctx context.Context, shortPath string) (string, error)
	CreateLink(ctx context.Context, longURL string) (string, error)
}

var (
	baseURL         = "http://localhost"
	long            = "https://jsonplaceholder.typicode.com/albums"
	short           = "n"
	existingLink, _ = nordpool.NewLink(1, long, time.Now())
)

func FuzzService_Create(f *testing.F) {
	// individual link test
	invididualLinkTests := []string{"abc123", "a1b2b3", "Ba_1bc-E"}
	for _, t := range invididualLinkTests {
		f.Add(t)
	}

	f.Fuzz(func(t *testing.T, in string) {
		inUrlEnc := base64.URLEncoding.EncodeToString([]byte(in))

		svc := NewNordpoolScraperService(newMockRepo(), baseURL)

		long := baseURL + "/" + inUrlEnc
		// make sure url len withing boundaries
		if nordpool.MaxLen < len(long) {
			long = long[:nordpool.MaxLen-len(baseURL)-1]
		}

		gotLink, err := svc.CreateLink(context.Background(), long)
		assert.NoError(t, err)
		// correct prefix
		assert.True(t, strings.HasPrefix(gotLink, fmt.Sprintf("%s/", baseURL)))
		// is alphanumeric
		short := strings.TrimPrefix(gotLink, fmt.Sprintf("%s/", baseURL))
		assert.NotEmpty(t, short)
		assert.True(t, govalidator.IsAlphanumeric(short))
		// redirects to original long uri
		gotInitialLong, err := svc.Redirect(context.Background(), short)
		assert.NoError(t, err)
		assert.Equal(t, long, gotInitialLong)
	})
}

func TestService_Create(t *testing.T) {
	// individual link test
	invididualLinkTests := []struct {
		name    string
		svc     linkService
		long    string
		wantErr bool
	}{
		{
			name:    "create link",
			svc:     NewNordpoolScraperService(newMockRepo(), baseURL),
			long:    long,
			wantErr: false,
		},
		{
			name:    "return existing link",
			svc:     NewNordpoolScraperService(newMockRepoWithSeed(existingLink), baseURL),
			long:    long,
			wantErr: false,
		},
	}
	for _, tt := range invididualLinkTests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.svc.CreateLink(context.Background(), tt.long)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}

	t.Run("multiple inserts", func(t *testing.T) {
		svc := NewNordpoolScraperService(newMockRepo(), baseURL)
		_, err := svc.CreateLink(context.Background(), long+"1")
		assert.NoError(t, err)
		_, err = svc.CreateLink(context.Background(), long+"2")
		assert.NoError(t, err)
		_, err = svc.CreateLink(context.Background(), long+"3")
		assert.NoError(t, err)
	})

	t.Run("return existing link", func(t *testing.T) {
		repoWithLink := newMockRepoWithSeed(existingLink)

		svc := NewNordpoolScraperService(repoWithLink, baseURL)
		l, err := svc.CreateLink(context.Background(), long)
		assert.NoError(t, err)
		assert.Equal(t, l, baseURL+"/"+existingLink.ShortPath())
	})
}

func TestService_Redirect(t *testing.T) {
	repoWithLink := newMockRepoWithSeed(existingLink)

	tests := []struct {
		name    string
		svc     linkService
		short   string
		wantErr bool
	}{
		{
			name:    "redirect",
			svc:     NewNordpoolScraperService(repoWithLink, baseURL),
			short:   short,
			wantErr: false,
		},
		{
			name:    "error doesn't exist",
			svc:     NewNordpoolScraperService(newMockRepo(), baseURL),
			short:   "asdasdad",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLink, err := tt.svc.Redirect(context.Background(), tt.short)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			// returned original link
			assert.Equal(t, gotLink, long)
		})
	}
}

package nordpool

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/alexgtn/go-linkshort/internal/domain/nordpool"
	ent "github.com/alexgtn/go-linkshort/tools/ent/codegen"
	"github.com/alexgtn/go-linkshort/tools/ent/codegen/nordpooldayahead"
)

// nordpoolRepo manages a sql-based link repo
type nordpoolRepo struct {
	client *ent.Client
}

func NewNordpoolSqliteRepo(c *ent.Client) *nordpoolRepo {
	return &nordpoolRepo{c}
}

func (n nordpoolRepo) InsertDayAheadValue(ctx context.Context, decimal decimal.Decimal, region nordpool.Region) (*nordpool.DayAheadValue, error) {
	// TODO implement me
}

func (n nordpoolRepo) GetDayAheadValue(ctx context.Context, date time.Time, region nordpool.Region) (*nordpool.DayAheadValue, error) {
	l, err := n.client.NordPoolDayAhead.
		Query().
		Where(nordpooldayahead.DateEQ(date.String())).
		Where(nordpooldayahead.RegionEQ(string(region))). // have to double check these two where haha
		Only(ctx)
	if err != nil {
		switch err.(type) {
		case *ent.NotFoundError:
			return nil, errors.Wrapf(pkg.ErrNotFound, "could not find)
		case *ent.NotSingularError:
			return nil, errors.Wrapf(pkg.ErrNotSingular, "more than one record")
		default:
			return nil, errors.Wrapf(err, "could not get")
		}
	}

	dayAheadValue, err := nordpool.NewDayAheadValue(..)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create day ahead value %s")
	}

	// construct this via constructor
	return dayAheadValue, nil
}

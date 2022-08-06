package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/martian/log"
	"github.com/shopspring/decimal"

	"github.com/alexgtn/go-linkshort/internal/domain/nordpool"
)

type nordpoolRepo interface {
	InsertDayAheadValue(ctx context.Context, decimal decimal.Decimal, region nordpool.Region) (*nordpool.DayAheadValue, error)
	GetDayAheadValue(ctx context.Context, date time.Time, region nordpool.Region) (*nordpool.DayAheadValue, error)
}

type service struct {
	ctx              context.Context
	nordpoolRepo     nordpoolRepo
	scrapingInterval *time.Ticker
	isDone           chan bool

	nordPoolURI string

	httpClient *resty.Client
}

func (s service) Stop() {
	s.scrapingInterval.Stop()
	s.isDone <- true
}

func (s service) Start() {
	for t := range s.scrapingInterval.C {
		fmt.Println(fmt.Sprintf("scraping nordpool stated at %s", t))

		for {
			select {
			case <-s.isDone:
				fmt.Println("done nordpool scraping")
				return
			case t := <-s.scrapingInterval.C:
				err := s.scrapeNordpool(s.ctx)
				if err != nil {
					log.Errorf(fmt.Sprintf("error in the scraping loop at %s, %v", t, err))
				}
			}
		}
	}
}

func NewNordpoolScraperService(
	ctx context.Context,
	linkRepo nordpoolRepo,
	scrapingInterval *time.Ticker,
	isDone chan bool,

	nordPoolURI string,

	httpClient *resty.Client) *service {
	return &service{
		// ...
	}
}

// Redirect returns the long URL provided a short path
func (s *service) scrapeNordpool(ctx context.Context) error {
	resty.New()
	//nordpoolPageResp, err := s.httpClient.R().Get(s.nordPoolURI)
	//if err != nil {
	//	return errors.Wrapf(err, "failed to get nordpool page via http")
	//}

	// this above won't work, content is returned dynamically .. dammit your Angular...

	// we will have an instance of puppeteer/cypress/or smth more golangy example https://github.com/mafredri/cdp
	// running in a separate docker instance
	// the problem that the content is dynamic
	// we need to wait until CSS selector is loaded (visible)
	// also handle connection issues which I've personally experienced

	// 1. select value
	// 2. check if not already stored for the current date, and insert (scrapers timings can be wonky)
	// fin
}

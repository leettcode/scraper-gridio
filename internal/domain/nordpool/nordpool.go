package nordpool

import (
	"time"

	"github.com/shopspring/decimal"
)

const MaxLen = 2048

type Region string

var FI Region = "FI"

// DayAheadValue is a link
type DayAheadValue struct {
	id int
	// eg. platform-base-url/shortPath
	date time.Time
	// original long uri that was shortened
	value     decimal.Decimal
	region    Region
	createdAt time.Time
}

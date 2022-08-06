// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// NordPoolDayAheadsColumns holds the columns for the "nord_pool_day_aheads" table.
	NordPoolDayAheadsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "date", Type: field.TypeString},
		{Name: "value", Type: field.TypeFloat64},
		{Name: "region", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
	}
	// NordPoolDayAheadsTable holds the schema information for the "nord_pool_day_aheads" table.
	NordPoolDayAheadsTable = &schema.Table{
		Name:       "nord_pool_day_aheads",
		Columns:    NordPoolDayAheadsColumns,
		PrimaryKey: []*schema.Column{NordPoolDayAheadsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		NordPoolDayAheadsTable,
	}
)

func init() {
}

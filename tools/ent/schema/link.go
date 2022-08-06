package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// NordPoolDayAhead holds the schema definition for the NordPoolDayAhead entity.
type NordPoolDayAhead struct {
	ent.Schema
}

// Fields of the NordPoolDayAhead.
func (NordPoolDayAhead) Fields() []ent.Field {
	return []ent.Field{
		// By default it creates a unique ID int
		field.String("date").NotEmpty(),
		field.Float("value").Positive(),
		field.String("region").NotEmpty(),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

// Edges of the NordPoolDayAhead.
func (NordPoolDayAhead) Edges() []ent.Edge {
	return nil
}

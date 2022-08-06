// Code generated by entc, DO NOT EDIT.

package codegen

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/alexgtn/go-linkshort/tools/ent/codegen/nordpooldayahead"
	"github.com/alexgtn/go-linkshort/tools/ent/codegen/predicate"
)

// NordPoolDayAheadUpdate is the builder for updating NordPoolDayAhead entities.
type NordPoolDayAheadUpdate struct {
	config
	hooks    []Hook
	mutation *NordPoolDayAheadMutation
}

// Where appends a list predicates to the NordPoolDayAheadUpdate builder.
func (npdau *NordPoolDayAheadUpdate) Where(ps ...predicate.NordPoolDayAhead) *NordPoolDayAheadUpdate {
	npdau.mutation.Where(ps...)
	return npdau
}

// SetDate sets the "date" field.
func (npdau *NordPoolDayAheadUpdate) SetDate(s string) *NordPoolDayAheadUpdate {
	npdau.mutation.SetDate(s)
	return npdau
}

// SetValue sets the "value" field.
func (npdau *NordPoolDayAheadUpdate) SetValue(f float64) *NordPoolDayAheadUpdate {
	npdau.mutation.ResetValue()
	npdau.mutation.SetValue(f)
	return npdau
}

// AddValue adds f to the "value" field.
func (npdau *NordPoolDayAheadUpdate) AddValue(f float64) *NordPoolDayAheadUpdate {
	npdau.mutation.AddValue(f)
	return npdau
}

// SetRegion sets the "region" field.
func (npdau *NordPoolDayAheadUpdate) SetRegion(s string) *NordPoolDayAheadUpdate {
	npdau.mutation.SetRegion(s)
	return npdau
}

// Mutation returns the NordPoolDayAheadMutation object of the builder.
func (npdau *NordPoolDayAheadUpdate) Mutation() *NordPoolDayAheadMutation {
	return npdau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (npdau *NordPoolDayAheadUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(npdau.hooks) == 0 {
		if err = npdau.check(); err != nil {
			return 0, err
		}
		affected, err = npdau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NordPoolDayAheadMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = npdau.check(); err != nil {
				return 0, err
			}
			npdau.mutation = mutation
			affected, err = npdau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(npdau.hooks) - 1; i >= 0; i-- {
			if npdau.hooks[i] == nil {
				return 0, fmt.Errorf("codegen: uninitialized hook (forgotten import codegen/runtime?)")
			}
			mut = npdau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, npdau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (npdau *NordPoolDayAheadUpdate) SaveX(ctx context.Context) int {
	affected, err := npdau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (npdau *NordPoolDayAheadUpdate) Exec(ctx context.Context) error {
	_, err := npdau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (npdau *NordPoolDayAheadUpdate) ExecX(ctx context.Context) {
	if err := npdau.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (npdau *NordPoolDayAheadUpdate) check() error {
	if v, ok := npdau.mutation.Date(); ok {
		if err := nordpooldayahead.DateValidator(v); err != nil {
			return &ValidationError{Name: "date", err: fmt.Errorf(`codegen: validator failed for field "NordPoolDayAhead.date": %w`, err)}
		}
	}
	if v, ok := npdau.mutation.Value(); ok {
		if err := nordpooldayahead.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`codegen: validator failed for field "NordPoolDayAhead.value": %w`, err)}
		}
	}
	if v, ok := npdau.mutation.Region(); ok {
		if err := nordpooldayahead.RegionValidator(v); err != nil {
			return &ValidationError{Name: "region", err: fmt.Errorf(`codegen: validator failed for field "NordPoolDayAhead.region": %w`, err)}
		}
	}
	return nil
}

func (npdau *NordPoolDayAheadUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   nordpooldayahead.Table,
			Columns: nordpooldayahead.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nordpooldayahead.FieldID,
			},
		},
	}
	if ps := npdau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := npdau.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nordpooldayahead.FieldDate,
		})
	}
	if value, ok := npdau.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: nordpooldayahead.FieldValue,
		})
	}
	if value, ok := npdau.mutation.AddedValue(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: nordpooldayahead.FieldValue,
		})
	}
	if value, ok := npdau.mutation.Region(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nordpooldayahead.FieldRegion,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, npdau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nordpooldayahead.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// NordPoolDayAheadUpdateOne is the builder for updating a single NordPoolDayAhead entity.
type NordPoolDayAheadUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NordPoolDayAheadMutation
}

// SetDate sets the "date" field.
func (npdauo *NordPoolDayAheadUpdateOne) SetDate(s string) *NordPoolDayAheadUpdateOne {
	npdauo.mutation.SetDate(s)
	return npdauo
}

// SetValue sets the "value" field.
func (npdauo *NordPoolDayAheadUpdateOne) SetValue(f float64) *NordPoolDayAheadUpdateOne {
	npdauo.mutation.ResetValue()
	npdauo.mutation.SetValue(f)
	return npdauo
}

// AddValue adds f to the "value" field.
func (npdauo *NordPoolDayAheadUpdateOne) AddValue(f float64) *NordPoolDayAheadUpdateOne {
	npdauo.mutation.AddValue(f)
	return npdauo
}

// SetRegion sets the "region" field.
func (npdauo *NordPoolDayAheadUpdateOne) SetRegion(s string) *NordPoolDayAheadUpdateOne {
	npdauo.mutation.SetRegion(s)
	return npdauo
}

// Mutation returns the NordPoolDayAheadMutation object of the builder.
func (npdauo *NordPoolDayAheadUpdateOne) Mutation() *NordPoolDayAheadMutation {
	return npdauo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (npdauo *NordPoolDayAheadUpdateOne) Select(field string, fields ...string) *NordPoolDayAheadUpdateOne {
	npdauo.fields = append([]string{field}, fields...)
	return npdauo
}

// Save executes the query and returns the updated NordPoolDayAhead entity.
func (npdauo *NordPoolDayAheadUpdateOne) Save(ctx context.Context) (*NordPoolDayAhead, error) {
	var (
		err  error
		node *NordPoolDayAhead
	)
	if len(npdauo.hooks) == 0 {
		if err = npdauo.check(); err != nil {
			return nil, err
		}
		node, err = npdauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NordPoolDayAheadMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = npdauo.check(); err != nil {
				return nil, err
			}
			npdauo.mutation = mutation
			node, err = npdauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(npdauo.hooks) - 1; i >= 0; i-- {
			if npdauo.hooks[i] == nil {
				return nil, fmt.Errorf("codegen: uninitialized hook (forgotten import codegen/runtime?)")
			}
			mut = npdauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, npdauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (npdauo *NordPoolDayAheadUpdateOne) SaveX(ctx context.Context) *NordPoolDayAhead {
	node, err := npdauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (npdauo *NordPoolDayAheadUpdateOne) Exec(ctx context.Context) error {
	_, err := npdauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (npdauo *NordPoolDayAheadUpdateOne) ExecX(ctx context.Context) {
	if err := npdauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (npdauo *NordPoolDayAheadUpdateOne) check() error {
	if v, ok := npdauo.mutation.Date(); ok {
		if err := nordpooldayahead.DateValidator(v); err != nil {
			return &ValidationError{Name: "date", err: fmt.Errorf(`codegen: validator failed for field "NordPoolDayAhead.date": %w`, err)}
		}
	}
	if v, ok := npdauo.mutation.Value(); ok {
		if err := nordpooldayahead.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`codegen: validator failed for field "NordPoolDayAhead.value": %w`, err)}
		}
	}
	if v, ok := npdauo.mutation.Region(); ok {
		if err := nordpooldayahead.RegionValidator(v); err != nil {
			return &ValidationError{Name: "region", err: fmt.Errorf(`codegen: validator failed for field "NordPoolDayAhead.region": %w`, err)}
		}
	}
	return nil
}

func (npdauo *NordPoolDayAheadUpdateOne) sqlSave(ctx context.Context) (_node *NordPoolDayAhead, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   nordpooldayahead.Table,
			Columns: nordpooldayahead.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nordpooldayahead.FieldID,
			},
		},
	}
	id, ok := npdauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`codegen: missing "NordPoolDayAhead.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := npdauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, nordpooldayahead.FieldID)
		for _, f := range fields {
			if !nordpooldayahead.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("codegen: invalid field %q for query", f)}
			}
			if f != nordpooldayahead.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := npdauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := npdauo.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nordpooldayahead.FieldDate,
		})
	}
	if value, ok := npdauo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: nordpooldayahead.FieldValue,
		})
	}
	if value, ok := npdauo.mutation.AddedValue(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: nordpooldayahead.FieldValue,
		})
	}
	if value, ok := npdauo.mutation.Region(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nordpooldayahead.FieldRegion,
		})
	}
	_node = &NordPoolDayAhead{config: npdauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, npdauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nordpooldayahead.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
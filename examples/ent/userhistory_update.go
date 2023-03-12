// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/frisbm/enthistory"
	"github.com/frisbm/enthistory/examples/ent/predicate"
	"github.com/frisbm/enthistory/examples/ent/userhistory"
)

// UserHistoryUpdate is the builder for updating UserHistory entities.
type UserHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *UserHistoryMutation
}

// Where appends a list predicates to the UserHistoryUpdate builder.
func (uhu *UserHistoryUpdate) Where(ps ...predicate.UserHistory) *UserHistoryUpdate {
	uhu.mutation.Where(ps...)
	return uhu
}

// SetHistoryTime sets the "history_time" field.
func (uhu *UserHistoryUpdate) SetHistoryTime(t time.Time) *UserHistoryUpdate {
	uhu.mutation.SetHistoryTime(t)
	return uhu
}

// SetRef sets the "ref" field.
func (uhu *UserHistoryUpdate) SetRef(i int) *UserHistoryUpdate {
	uhu.mutation.ResetRef()
	uhu.mutation.SetRef(i)
	return uhu
}

// SetNillableRef sets the "ref" field if the given value is not nil.
func (uhu *UserHistoryUpdate) SetNillableRef(i *int) *UserHistoryUpdate {
	if i != nil {
		uhu.SetRef(*i)
	}
	return uhu
}

// AddRef adds i to the "ref" field.
func (uhu *UserHistoryUpdate) AddRef(i int) *UserHistoryUpdate {
	uhu.mutation.AddRef(i)
	return uhu
}

// ClearRef clears the value of the "ref" field.
func (uhu *UserHistoryUpdate) ClearRef() *UserHistoryUpdate {
	uhu.mutation.ClearRef()
	return uhu
}

// SetUpdatedBy sets the "updated_by" field.
func (uhu *UserHistoryUpdate) SetUpdatedBy(i int) *UserHistoryUpdate {
	uhu.mutation.ResetUpdatedBy()
	uhu.mutation.SetUpdatedBy(i)
	return uhu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (uhu *UserHistoryUpdate) SetNillableUpdatedBy(i *int) *UserHistoryUpdate {
	if i != nil {
		uhu.SetUpdatedBy(*i)
	}
	return uhu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (uhu *UserHistoryUpdate) AddUpdatedBy(i int) *UserHistoryUpdate {
	uhu.mutation.AddUpdatedBy(i)
	return uhu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (uhu *UserHistoryUpdate) ClearUpdatedBy() *UserHistoryUpdate {
	uhu.mutation.ClearUpdatedBy()
	return uhu
}

// SetOperation sets the "operation" field.
func (uhu *UserHistoryUpdate) SetOperation(et enthistory.OpType) *UserHistoryUpdate {
	uhu.mutation.SetOperation(et)
	return uhu
}

// SetCreatedAt sets the "created_at" field.
func (uhu *UserHistoryUpdate) SetCreatedAt(t time.Time) *UserHistoryUpdate {
	uhu.mutation.SetCreatedAt(t)
	return uhu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uhu *UserHistoryUpdate) SetNillableCreatedAt(t *time.Time) *UserHistoryUpdate {
	if t != nil {
		uhu.SetCreatedAt(*t)
	}
	return uhu
}

// SetUpdatedAt sets the "updated_at" field.
func (uhu *UserHistoryUpdate) SetUpdatedAt(t time.Time) *UserHistoryUpdate {
	uhu.mutation.SetUpdatedAt(t)
	return uhu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uhu *UserHistoryUpdate) SetNillableUpdatedAt(t *time.Time) *UserHistoryUpdate {
	if t != nil {
		uhu.SetUpdatedAt(*t)
	}
	return uhu
}

// SetAge sets the "age" field.
func (uhu *UserHistoryUpdate) SetAge(i int) *UserHistoryUpdate {
	uhu.mutation.ResetAge()
	uhu.mutation.SetAge(i)
	return uhu
}

// AddAge adds i to the "age" field.
func (uhu *UserHistoryUpdate) AddAge(i int) *UserHistoryUpdate {
	uhu.mutation.AddAge(i)
	return uhu
}

// SetName sets the "name" field.
func (uhu *UserHistoryUpdate) SetName(s string) *UserHistoryUpdate {
	uhu.mutation.SetName(s)
	return uhu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uhu *UserHistoryUpdate) SetNillableName(s *string) *UserHistoryUpdate {
	if s != nil {
		uhu.SetName(*s)
	}
	return uhu
}

// Mutation returns the UserHistoryMutation object of the builder.
func (uhu *UserHistoryUpdate) Mutation() *UserHistoryMutation {
	return uhu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uhu *UserHistoryUpdate) Save(ctx context.Context) (int, error) {
	uhu.defaults()
	return withHooks[int, UserHistoryMutation](ctx, uhu.sqlSave, uhu.mutation, uhu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uhu *UserHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := uhu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uhu *UserHistoryUpdate) Exec(ctx context.Context) error {
	_, err := uhu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uhu *UserHistoryUpdate) ExecX(ctx context.Context) {
	if err := uhu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uhu *UserHistoryUpdate) defaults() {
	if _, ok := uhu.mutation.HistoryTime(); !ok {
		v := userhistory.UpdateDefaultHistoryTime()
		uhu.mutation.SetHistoryTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uhu *UserHistoryUpdate) check() error {
	if v, ok := uhu.mutation.Operation(); ok {
		if err := userhistory.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`ent: validator failed for field "UserHistory.operation": %w`, err)}
		}
	}
	if v, ok := uhu.mutation.Age(); ok {
		if err := userhistory.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "UserHistory.age": %w`, err)}
		}
	}
	return nil
}

func (uhu *UserHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uhu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(userhistory.Table, userhistory.Columns, sqlgraph.NewFieldSpec(userhistory.FieldID, field.TypeInt))
	if ps := uhu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uhu.mutation.HistoryTime(); ok {
		_spec.SetField(userhistory.FieldHistoryTime, field.TypeTime, value)
	}
	if value, ok := uhu.mutation.Ref(); ok {
		_spec.SetField(userhistory.FieldRef, field.TypeInt, value)
	}
	if value, ok := uhu.mutation.AddedRef(); ok {
		_spec.AddField(userhistory.FieldRef, field.TypeInt, value)
	}
	if uhu.mutation.RefCleared() {
		_spec.ClearField(userhistory.FieldRef, field.TypeInt)
	}
	if value, ok := uhu.mutation.UpdatedBy(); ok {
		_spec.SetField(userhistory.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := uhu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(userhistory.FieldUpdatedBy, field.TypeInt, value)
	}
	if uhu.mutation.UpdatedByCleared() {
		_spec.ClearField(userhistory.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := uhu.mutation.Operation(); ok {
		_spec.SetField(userhistory.FieldOperation, field.TypeEnum, value)
	}
	if value, ok := uhu.mutation.CreatedAt(); ok {
		_spec.SetField(userhistory.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uhu.mutation.UpdatedAt(); ok {
		_spec.SetField(userhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uhu.mutation.Age(); ok {
		_spec.SetField(userhistory.FieldAge, field.TypeInt, value)
	}
	if value, ok := uhu.mutation.AddedAge(); ok {
		_spec.AddField(userhistory.FieldAge, field.TypeInt, value)
	}
	if value, ok := uhu.mutation.Name(); ok {
		_spec.SetField(userhistory.FieldName, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uhu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uhu.mutation.done = true
	return n, nil
}

// UserHistoryUpdateOne is the builder for updating a single UserHistory entity.
type UserHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserHistoryMutation
}

// SetHistoryTime sets the "history_time" field.
func (uhuo *UserHistoryUpdateOne) SetHistoryTime(t time.Time) *UserHistoryUpdateOne {
	uhuo.mutation.SetHistoryTime(t)
	return uhuo
}

// SetRef sets the "ref" field.
func (uhuo *UserHistoryUpdateOne) SetRef(i int) *UserHistoryUpdateOne {
	uhuo.mutation.ResetRef()
	uhuo.mutation.SetRef(i)
	return uhuo
}

// SetNillableRef sets the "ref" field if the given value is not nil.
func (uhuo *UserHistoryUpdateOne) SetNillableRef(i *int) *UserHistoryUpdateOne {
	if i != nil {
		uhuo.SetRef(*i)
	}
	return uhuo
}

// AddRef adds i to the "ref" field.
func (uhuo *UserHistoryUpdateOne) AddRef(i int) *UserHistoryUpdateOne {
	uhuo.mutation.AddRef(i)
	return uhuo
}

// ClearRef clears the value of the "ref" field.
func (uhuo *UserHistoryUpdateOne) ClearRef() *UserHistoryUpdateOne {
	uhuo.mutation.ClearRef()
	return uhuo
}

// SetUpdatedBy sets the "updated_by" field.
func (uhuo *UserHistoryUpdateOne) SetUpdatedBy(i int) *UserHistoryUpdateOne {
	uhuo.mutation.ResetUpdatedBy()
	uhuo.mutation.SetUpdatedBy(i)
	return uhuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (uhuo *UserHistoryUpdateOne) SetNillableUpdatedBy(i *int) *UserHistoryUpdateOne {
	if i != nil {
		uhuo.SetUpdatedBy(*i)
	}
	return uhuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (uhuo *UserHistoryUpdateOne) AddUpdatedBy(i int) *UserHistoryUpdateOne {
	uhuo.mutation.AddUpdatedBy(i)
	return uhuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (uhuo *UserHistoryUpdateOne) ClearUpdatedBy() *UserHistoryUpdateOne {
	uhuo.mutation.ClearUpdatedBy()
	return uhuo
}

// SetOperation sets the "operation" field.
func (uhuo *UserHistoryUpdateOne) SetOperation(et enthistory.OpType) *UserHistoryUpdateOne {
	uhuo.mutation.SetOperation(et)
	return uhuo
}

// SetCreatedAt sets the "created_at" field.
func (uhuo *UserHistoryUpdateOne) SetCreatedAt(t time.Time) *UserHistoryUpdateOne {
	uhuo.mutation.SetCreatedAt(t)
	return uhuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uhuo *UserHistoryUpdateOne) SetNillableCreatedAt(t *time.Time) *UserHistoryUpdateOne {
	if t != nil {
		uhuo.SetCreatedAt(*t)
	}
	return uhuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uhuo *UserHistoryUpdateOne) SetUpdatedAt(t time.Time) *UserHistoryUpdateOne {
	uhuo.mutation.SetUpdatedAt(t)
	return uhuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uhuo *UserHistoryUpdateOne) SetNillableUpdatedAt(t *time.Time) *UserHistoryUpdateOne {
	if t != nil {
		uhuo.SetUpdatedAt(*t)
	}
	return uhuo
}

// SetAge sets the "age" field.
func (uhuo *UserHistoryUpdateOne) SetAge(i int) *UserHistoryUpdateOne {
	uhuo.mutation.ResetAge()
	uhuo.mutation.SetAge(i)
	return uhuo
}

// AddAge adds i to the "age" field.
func (uhuo *UserHistoryUpdateOne) AddAge(i int) *UserHistoryUpdateOne {
	uhuo.mutation.AddAge(i)
	return uhuo
}

// SetName sets the "name" field.
func (uhuo *UserHistoryUpdateOne) SetName(s string) *UserHistoryUpdateOne {
	uhuo.mutation.SetName(s)
	return uhuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uhuo *UserHistoryUpdateOne) SetNillableName(s *string) *UserHistoryUpdateOne {
	if s != nil {
		uhuo.SetName(*s)
	}
	return uhuo
}

// Mutation returns the UserHistoryMutation object of the builder.
func (uhuo *UserHistoryUpdateOne) Mutation() *UserHistoryMutation {
	return uhuo.mutation
}

// Where appends a list predicates to the UserHistoryUpdate builder.
func (uhuo *UserHistoryUpdateOne) Where(ps ...predicate.UserHistory) *UserHistoryUpdateOne {
	uhuo.mutation.Where(ps...)
	return uhuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uhuo *UserHistoryUpdateOne) Select(field string, fields ...string) *UserHistoryUpdateOne {
	uhuo.fields = append([]string{field}, fields...)
	return uhuo
}

// Save executes the query and returns the updated UserHistory entity.
func (uhuo *UserHistoryUpdateOne) Save(ctx context.Context) (*UserHistory, error) {
	uhuo.defaults()
	return withHooks[*UserHistory, UserHistoryMutation](ctx, uhuo.sqlSave, uhuo.mutation, uhuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uhuo *UserHistoryUpdateOne) SaveX(ctx context.Context) *UserHistory {
	node, err := uhuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uhuo *UserHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := uhuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uhuo *UserHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := uhuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uhuo *UserHistoryUpdateOne) defaults() {
	if _, ok := uhuo.mutation.HistoryTime(); !ok {
		v := userhistory.UpdateDefaultHistoryTime()
		uhuo.mutation.SetHistoryTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uhuo *UserHistoryUpdateOne) check() error {
	if v, ok := uhuo.mutation.Operation(); ok {
		if err := userhistory.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`ent: validator failed for field "UserHistory.operation": %w`, err)}
		}
	}
	if v, ok := uhuo.mutation.Age(); ok {
		if err := userhistory.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "UserHistory.age": %w`, err)}
		}
	}
	return nil
}

func (uhuo *UserHistoryUpdateOne) sqlSave(ctx context.Context) (_node *UserHistory, err error) {
	if err := uhuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(userhistory.Table, userhistory.Columns, sqlgraph.NewFieldSpec(userhistory.FieldID, field.TypeInt))
	id, ok := uhuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uhuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userhistory.FieldID)
		for _, f := range fields {
			if !userhistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uhuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uhuo.mutation.HistoryTime(); ok {
		_spec.SetField(userhistory.FieldHistoryTime, field.TypeTime, value)
	}
	if value, ok := uhuo.mutation.Ref(); ok {
		_spec.SetField(userhistory.FieldRef, field.TypeInt, value)
	}
	if value, ok := uhuo.mutation.AddedRef(); ok {
		_spec.AddField(userhistory.FieldRef, field.TypeInt, value)
	}
	if uhuo.mutation.RefCleared() {
		_spec.ClearField(userhistory.FieldRef, field.TypeInt)
	}
	if value, ok := uhuo.mutation.UpdatedBy(); ok {
		_spec.SetField(userhistory.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := uhuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(userhistory.FieldUpdatedBy, field.TypeInt, value)
	}
	if uhuo.mutation.UpdatedByCleared() {
		_spec.ClearField(userhistory.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := uhuo.mutation.Operation(); ok {
		_spec.SetField(userhistory.FieldOperation, field.TypeEnum, value)
	}
	if value, ok := uhuo.mutation.CreatedAt(); ok {
		_spec.SetField(userhistory.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uhuo.mutation.UpdatedAt(); ok {
		_spec.SetField(userhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uhuo.mutation.Age(); ok {
		_spec.SetField(userhistory.FieldAge, field.TypeInt, value)
	}
	if value, ok := uhuo.mutation.AddedAge(); ok {
		_spec.AddField(userhistory.FieldAge, field.TypeInt, value)
	}
	if value, ok := uhuo.mutation.Name(); ok {
		_spec.SetField(userhistory.FieldName, field.TypeString, value)
	}
	_node = &UserHistory{config: uhuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uhuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uhuo.mutation.done = true
	return _node, nil
}

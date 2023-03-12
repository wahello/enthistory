// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"

	"github.com/frisbm/enthistory"
	"github.com/frisbm/enthistory/examples/ent/userhistory"
)

// UserHistory is the model entity for the UserHistory schema.
type UserHistory struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// HistoryTime holds the value of the "history_time" field.
	HistoryTime time.Time `json:"history_time,omitempty"`
	// Ref holds the value of the "ref" field.
	Ref int `json:"ref,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy int `json:"updated_by,omitempty"`
	// Operation holds the value of the "operation" field.
	Operation enthistory.OpType `json:"operation,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userhistory.FieldID, userhistory.FieldRef, userhistory.FieldUpdatedBy, userhistory.FieldAge:
			values[i] = new(sql.NullInt64)
		case userhistory.FieldOperation, userhistory.FieldName:
			values[i] = new(sql.NullString)
		case userhistory.FieldHistoryTime, userhistory.FieldCreatedAt, userhistory.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UserHistory", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserHistory fields.
func (uh *UserHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userhistory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			uh.ID = int(value.Int64)
		case userhistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				uh.HistoryTime = value.Time
			}
		case userhistory.FieldRef:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value.Valid {
				uh.Ref = int(value.Int64)
			}
		case userhistory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				uh.UpdatedBy = int(value.Int64)
			}
		case userhistory.FieldOperation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value.Valid {
				uh.Operation = enthistory.OpType(value.String)
			}
		case userhistory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				uh.CreatedAt = value.Time
			}
		case userhistory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				uh.UpdatedAt = value.Time
			}
		case userhistory.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				uh.Age = int(value.Int64)
			}
		case userhistory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				uh.Name = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this UserHistory.
// Note that you need to call UserHistory.Unwrap() before calling this method if this UserHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (uh *UserHistory) Update() *UserHistoryUpdateOne {
	return NewUserHistoryClient(uh.config).UpdateOne(uh)
}

// Unwrap unwraps the UserHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uh *UserHistory) Unwrap() *UserHistory {
	_tx, ok := uh.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserHistory is not a transactional entity")
	}
	uh.config.driver = _tx.drv
	return uh
}

// String implements the fmt.Stringer.
func (uh *UserHistory) String() string {
	var builder strings.Builder
	builder.WriteString("UserHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", uh.ID))
	builder.WriteString("history_time=")
	builder.WriteString(uh.HistoryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(fmt.Sprintf("%v", uh.Ref))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", uh.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(fmt.Sprintf("%v", uh.Operation))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(uh.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(uh.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("age=")
	builder.WriteString(fmt.Sprintf("%v", uh.Age))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(uh.Name)
	builder.WriteByte(')')
	return builder.String()
}

// UserHistories is a parsable slice of UserHistory.
type UserHistories []*UserHistory

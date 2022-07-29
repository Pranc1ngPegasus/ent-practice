// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Pranc1ngPegasus/ent-practice/ent/organization"
	"github.com/google/uuid"
)

// Organization is the model entity for the Organization schema.
type Organization struct {
	config `json:"-"`
	// ID of the ent.
	// ID
	ID uuid.UUID `json:"id,omitempty"`
	// 組織名
	Name string `json:"name,omitempty"`
	// 作成日
	CreatedAt time.Time `json:"created_at,omitempty"`
	// 更新日
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// 削除日
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Organization) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case organization.FieldName:
			values[i] = new(sql.NullString)
		case organization.FieldCreatedAt, organization.FieldUpdatedAt, organization.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case organization.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Organization", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Organization fields.
func (o *Organization) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case organization.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				o.ID = *value
			}
		case organization.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				o.Name = value.String
			}
		case organization.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				o.CreatedAt = value.Time
			}
		case organization.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				o.UpdatedAt = value.Time
			}
		case organization.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				o.DeletedAt = new(time.Time)
				*o.DeletedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Organization.
// Note that you need to call Organization.Unwrap() before calling this method if this Organization
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Organization) Update() *OrganizationUpdateOne {
	return (&OrganizationClient{config: o.config}).UpdateOne(o)
}

// Unwrap unwraps the Organization entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Organization) Unwrap() *Organization {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Organization is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Organization) String() string {
	var builder strings.Builder
	builder.WriteString("Organization(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("name=")
	builder.WriteString(o.Name)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(o.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(o.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := o.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Organizations is a parsable slice of Organization.
type Organizations []*Organization

func (o Organizations) config(cfg config) {
	for _i := range o {
		o[_i].config = cfg
	}
}

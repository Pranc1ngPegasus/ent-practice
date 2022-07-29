// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"time"

	"github.com/go-faster/jx"
	"github.com/google/uuid"
)

// Ref: #/components/schemas/AccountCreate
type AccountCreate struct {
	ID        uuid.UUID   "json:\"id\""
	Name      string      "json:\"name\""
	CreatedAt time.Time   "json:\"created_at\""
	UpdatedAt time.Time   "json:\"updated_at\""
	DeletedAt OptDateTime "json:\"deleted_at\""
}

func (*AccountCreate) createAccountRes() {}

// Ref: #/components/schemas/AccountList
type AccountList struct {
	ID        uuid.UUID   "json:\"id\""
	Name      string      "json:\"name\""
	CreatedAt time.Time   "json:\"created_at\""
	UpdatedAt time.Time   "json:\"updated_at\""
	DeletedAt OptDateTime "json:\"deleted_at\""
}

// Ref: #/components/schemas/AccountRead
type AccountRead struct {
	ID        uuid.UUID   "json:\"id\""
	Name      string      "json:\"name\""
	CreatedAt time.Time   "json:\"created_at\""
	UpdatedAt time.Time   "json:\"updated_at\""
	DeletedAt OptDateTime "json:\"deleted_at\""
}

func (*AccountRead) readAccountRes() {}

// Ref: #/components/schemas/AccountUpdate
type AccountUpdate struct {
	ID        uuid.UUID   "json:\"id\""
	Name      string      "json:\"name\""
	CreatedAt time.Time   "json:\"created_at\""
	UpdatedAt time.Time   "json:\"updated_at\""
	DeletedAt OptDateTime "json:\"deleted_at\""
}

func (*AccountUpdate) updateAccountRes() {}

type CreateAccountReq struct {
	Name      string      "json:\"name\""
	CreatedAt time.Time   "json:\"created_at\""
	UpdatedAt time.Time   "json:\"updated_at\""
	DeletedAt OptDateTime "json:\"deleted_at\""
}

type CreateOrganizationReq struct {
	Name      string      "json:\"name\""
	CreatedAt time.Time   "json:\"created_at\""
	UpdatedAt time.Time   "json:\"updated_at\""
	DeletedAt OptDateTime "json:\"deleted_at\""
}

// DeleteAccountNoContent is response for DeleteAccount operation.
type DeleteAccountNoContent struct{}

func (*DeleteAccountNoContent) deleteAccountRes() {}

// DeleteOrganizationNoContent is response for DeleteOrganization operation.
type DeleteOrganizationNoContent struct{}

func (*DeleteOrganizationNoContent) deleteOrganizationRes() {}

type ListAccountOKApplicationJSON []AccountList

func (ListAccountOKApplicationJSON) listAccountRes() {}

type ListOrganizationOKApplicationJSON []OrganizationList

func (ListOrganizationOKApplicationJSON) listOrganizationRes() {}

// NewOptDateTime returns new OptDateTime with value set to v.
func NewOptDateTime(v time.Time) OptDateTime {
	return OptDateTime{
		Value: v,
		Set:   true,
	}
}

// OptDateTime is optional time.Time.
type OptDateTime struct {
	Value time.Time
	Set   bool
}

// IsSet returns true if OptDateTime was set.
func (o OptDateTime) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDateTime) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDateTime) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDateTime) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDateTime) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/OrganizationCreate
type OrganizationCreate struct {
	ID        uuid.UUID   "json:\"id\""
	Name      string      "json:\"name\""
	CreatedAt time.Time   "json:\"created_at\""
	UpdatedAt time.Time   "json:\"updated_at\""
	DeletedAt OptDateTime "json:\"deleted_at\""
}

func (*OrganizationCreate) createOrganizationRes() {}

// Ref: #/components/schemas/OrganizationList
type OrganizationList struct {
	ID        uuid.UUID   "json:\"id\""
	Name      string      "json:\"name\""
	CreatedAt time.Time   "json:\"created_at\""
	UpdatedAt time.Time   "json:\"updated_at\""
	DeletedAt OptDateTime "json:\"deleted_at\""
}

// Ref: #/components/schemas/OrganizationRead
type OrganizationRead struct {
	ID        uuid.UUID   "json:\"id\""
	Name      string      "json:\"name\""
	CreatedAt time.Time   "json:\"created_at\""
	UpdatedAt time.Time   "json:\"updated_at\""
	DeletedAt OptDateTime "json:\"deleted_at\""
}

func (*OrganizationRead) readOrganizationRes() {}

// Ref: #/components/schemas/OrganizationUpdate
type OrganizationUpdate struct {
	ID        uuid.UUID   "json:\"id\""
	Name      string      "json:\"name\""
	CreatedAt time.Time   "json:\"created_at\""
	UpdatedAt time.Time   "json:\"updated_at\""
	DeletedAt OptDateTime "json:\"deleted_at\""
}

func (*OrganizationUpdate) updateOrganizationRes() {}

type R400 struct {
	Code   int    "json:\"code\""
	Status string "json:\"status\""
	Errors jx.Raw "json:\"errors\""
}

func (*R400) createAccountRes()      {}
func (*R400) createOrganizationRes() {}
func (*R400) deleteAccountRes()      {}
func (*R400) deleteOrganizationRes() {}
func (*R400) listAccountRes()        {}
func (*R400) listOrganizationRes()   {}
func (*R400) readAccountRes()        {}
func (*R400) readOrganizationRes()   {}
func (*R400) updateAccountRes()      {}
func (*R400) updateOrganizationRes() {}

type R404 struct {
	Code   int    "json:\"code\""
	Status string "json:\"status\""
	Errors jx.Raw "json:\"errors\""
}

func (*R404) deleteAccountRes()      {}
func (*R404) deleteOrganizationRes() {}
func (*R404) listAccountRes()        {}
func (*R404) listOrganizationRes()   {}
func (*R404) readAccountRes()        {}
func (*R404) readOrganizationRes()   {}
func (*R404) updateAccountRes()      {}
func (*R404) updateOrganizationRes() {}

type R409 struct {
	Code   int    "json:\"code\""
	Status string "json:\"status\""
	Errors jx.Raw "json:\"errors\""
}

func (*R409) createAccountRes()      {}
func (*R409) createOrganizationRes() {}
func (*R409) deleteAccountRes()      {}
func (*R409) deleteOrganizationRes() {}
func (*R409) listAccountRes()        {}
func (*R409) listOrganizationRes()   {}
func (*R409) readAccountRes()        {}
func (*R409) readOrganizationRes()   {}
func (*R409) updateAccountRes()      {}
func (*R409) updateOrganizationRes() {}

type R500 struct {
	Code   int    "json:\"code\""
	Status string "json:\"status\""
	Errors jx.Raw "json:\"errors\""
}

func (*R500) createAccountRes()      {}
func (*R500) createOrganizationRes() {}
func (*R500) deleteAccountRes()      {}
func (*R500) deleteOrganizationRes() {}
func (*R500) listAccountRes()        {}
func (*R500) listOrganizationRes()   {}
func (*R500) readAccountRes()        {}
func (*R500) readOrganizationRes()   {}
func (*R500) updateAccountRes()      {}
func (*R500) updateOrganizationRes() {}

type UpdateAccountReq struct {
	Name      OptString   "json:\"name\""
	UpdatedAt OptDateTime "json:\"updated_at\""
	DeletedAt OptDateTime "json:\"deleted_at\""
}

type UpdateOrganizationReq struct {
	Name      OptString   "json:\"name\""
	UpdatedAt OptDateTime "json:\"updated_at\""
	DeletedAt OptDateTime "json:\"deleted_at\""
}

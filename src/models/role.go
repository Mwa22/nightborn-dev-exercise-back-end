package models

// swagger:enum RoleEnum
type RoleEnum string

const (
	Administrator RoleEnum = "Administrator"
	RegularUser   RoleEnum = "RegularUser"
)

// A role for an user
//
// swagger:model
type Role RoleEnum

func (r Role) IsValid() bool {
	return r == "Administrator" || r == "RegularUser"
}

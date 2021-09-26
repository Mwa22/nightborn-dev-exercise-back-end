package models

// A role for an user
//
// swagger:model
type Role int16

const (
	Administrator Role = iota
	RegularUser
)

func (r Role) IsValid() bool {
	return r >= Administrator && r <= RegularUser
}

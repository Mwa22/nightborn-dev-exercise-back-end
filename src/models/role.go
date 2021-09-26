package models

type Role int16

const (
	Administrator Role = iota
	RegularUser
)

func (r Role) IsValid() bool {
	return r >= Administrator && r <= RegularUser
}

package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

// An user of the application
//
// swagger:model
type User struct {
	gorm.Model

	// The firstname
	//
	// required: true
	FIRSTNAME string `json:"first_name" validate:"required"`

	// The lastname
	//
	// required: true
	LASTNAME string `json:"last_name" validate:"required"`

	// The email
	//
	// required: true
	// unique: true
	// example: user@provider.net
	EMAIL string `json:"email" validate:"required,email" gorm:"unique_index"`

	// The role
	//
	// required: true
	ROLE Role `json:"role" validate:"required,role"`

	// The password
	//
	// required: true
	PASSWORD string `json:"password" validate:"required"`
}

// An object to update a user
//
// swagger:model
type UpdateUserInput struct {
	// The firstname
	//
	// required: false
	FIRSTNAME string `json:"first_name"`

	// The lastname
	//
	// required: false
	LASTNAME string `json:"last_name"`

	// The email
	//
	// required: false
	// unique: true
	// example: user@provider.net
	EMAIL string `json:"email" validate:"email" gorm:"unique_index"`

	// The role
	//
	// required: false
	ROLE Role `json:"role" validate:"role"`

	// The password
	//
	// required: false
	PASSWORD string `json:"password"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	validate := validator.New()
	_ = validate.RegisterValidation("role", func(fl validator.FieldLevel) bool {
		return Role(fl.Field().String()).IsValid()
	})

	err = validate.Struct(u)
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	validate := validator.New()
	_ = validate.RegisterValidation("role", func(fl validator.FieldLevel) bool {
		return Role(fl.Field().String()).IsValid()
	})

	err = validate.Struct(u)
	return
}

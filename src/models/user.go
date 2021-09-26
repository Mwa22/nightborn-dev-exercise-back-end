package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FIRSTNAME string `json:"first_name" validate:"required"`
	LASTNAME  string `json:"last_name" validate:"required"`
	EMAIL     string `json:"email" validate:"required,email" gorm:"unique_index"`
	ROLE      *Role  `json:"role" validate:"required,role"`
	PASSWORD  string `json:"password" validate:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	validate := validator.New()
	_ = validate.RegisterValidation("role", func(fl validator.FieldLevel) bool {
		return Role(fl.Field().Int()).IsValid()
	})

	err = validate.Struct(u)
	return
}

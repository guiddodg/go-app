package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string
}

func NewUser(email, password string) (*User, error) {
	//todo: validations should be here and not in controller
	return &User{
		Email:    email,
		Password: password,
	}, nil
}

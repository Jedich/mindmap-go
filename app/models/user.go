package models

import (
	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	FirstName string  `json:"first_name,omitempty" query:"first_name"`
	LastName  string  `json:"last_name,omitempty" query:"last_name"`
	AccountID int     `json:"account_id" query:"account_id" gorm:"index"`
	Account   Account `gorm:"foreignKey:AccountID"`
}

type UserUpdate struct {
	FirstName string
	LastName  string
}

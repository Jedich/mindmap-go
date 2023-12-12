package models

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        int            `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// User model
type User struct {
	Model
	FirstName string  `json:"first_name,omitempty"`
	LastName  string  `json:"last_name,omitempty"`
	AccountID int     `json:"-" gorm:"index"`
	Account   Account `json:"account" gorm:"foreignKey:AccountID"`
}

type UserUpdate struct {
	FirstName string
	LastName  string
}

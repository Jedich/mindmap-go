package models

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// User model
type User struct {
	Model
	FirstName string  `json:"first_name,omitempty" query:"first_name"`
	LastName  string  `json:"last_name,omitempty" query:"last_name"`
	AccountID int     `json:"-" query:"account_id" gorm:"index"`
	Account   Account `json:"account" gorm:"foreignKey:AccountID"`
	Maps      []Map   `json:"maps,omitempty" gorm:"foreignKey:CreatorID"`
	Cards     []Card  `json:"-" gorm:"foreignKey:CreatorID"`
}

type UserUpdate struct {
	FirstName string
	LastName  string
}

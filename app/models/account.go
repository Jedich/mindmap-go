package models

import "gorm.io/gorm"

// Account model
type Account struct {
	gorm.Model
	Username     string `json:"username" query:"username" gorm:"index"`
	Email        string `json:"email" query:"email" gorm:"index"`
	Password     string `json:"-" query:"-"`
	PasswordHash string `json:"password_hash" query:"password_hash"`
}

type AccountUpdate struct {
	Username     string `json:"username" query:"username"`
	PasswordHash string `json:"password_hash" query:"password_hash"`
}

func (u *Account) Validate() error {
	return nil
}

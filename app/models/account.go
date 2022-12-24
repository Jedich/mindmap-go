package models

// Account model
type Account struct {
	Model        `json:"-"`
	Username     string `json:"username" gorm:"index"`
	Email        string `json:"email" gorm:"index"`
	PasswordHash []byte `json:"-"`
}

type AccountUpdate struct {
	Username     string `json:"username"`
	PasswordHash []byte `json:"password_hash"`
}

func (u *Account) Validate() error {
	return nil
}

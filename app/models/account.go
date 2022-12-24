package models

// Account model
type Account struct {
	Model        `json:"-"`
	Username     string `json:"username" query:"username" gorm:"index"`
	Email        string `json:"email" query:"email" gorm:"index"`
	PasswordHash []byte `json:"-" query:"password_hash"`
}

type AccountUpdate struct {
	Username     string `json:"username" query:"username"`
	PasswordHash []byte `json:"password_hash" query:"password_hash"`
}

func (u *Account) Validate() error {
	return nil
}

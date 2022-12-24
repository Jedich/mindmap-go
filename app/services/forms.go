package services

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type RegisterForm struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

func (u RegisterForm) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Username, validation.Required, validation.Length(5, 45)),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(8, 45)),
		validation.Field(&u.FirstName, validation.Length(5, 45)),
		validation.Field(&u.LastName, validation.Length(5, 45)),
	)
}

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u LoginForm) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(8, 45)),
	)
}

type MapForm struct {
	CreatorID   int    `json:"creator_id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"desc,omitempty"`
}

func (m MapForm) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Length(0, 45)),
		validation.Field(&m.CreatorID, validation.Required, validation.NotNil),
	)
}

type CardForm struct {
	Name      string   `json:"name"`
	Text      string   `json:"text_data"`
	Level     *int     `json:"level"`
	PositionY *float64 `json:"position_y"`
	ParentID  *int     `json:"parent_id"`
	CreatorID int      `json:"creator_id"`
	MapID     int      `json:"map_id"`
}

func (c CardForm) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Length(0, 255)),
		validation.Field(&c.Level, validation.NotNil),
		validation.Field(&c.PositionY, validation.NotNil),
		validation.Field(&c.CreatorID, validation.Required, validation.NotNil),
		validation.Field(&c.MapID, validation.Required, validation.NotNil),
	)
}

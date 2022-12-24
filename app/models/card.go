package models

import validation "github.com/go-ozzo/ozzo-validation"

// Card model
type Card struct {
	Model
	Name      string   `json:"name"`
	Text      string   `json:"text_data" gorm:"column:text_data"`
	Level     *int     `json:"level"`
	PositionY *float64 `json:"position_y"`
	ParentID  *int     `json:"parent_id"`
	CreatorID int      `json:"creator_id"`
	MapID     int      `json:"map_id"`
	Parent    *Card    `json:"-" gorm:"foreignKey:ParentID"`
}

type CardUpdate struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Text      string   `json:"text_data"`
	Level     *int     `json:"level"`
	PositionY *float64 `json:"position_y"`
	ParentID  *int     `json:"parent_id"`
}

func (c CardUpdate) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Length(0, 255)),
		validation.Field(&c.ID, validation.Required),
	)
}

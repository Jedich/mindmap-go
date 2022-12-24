package models

// Card model
type Card struct {
	Model
	Name      string `json:"name" query:"name"`
	Text      string `json:"text_data" gorm:"column:text_data"`
	ParentID  *int   `json:"parent_id" query:"parent_id"`
	CreatorID int    `json:"creator_id" query:"creator_id"`
	MapID     int    `json:"map_id" query:"map_id"`
	Parent    *Card  `gorm:"foreignKey:ParentID"`
}

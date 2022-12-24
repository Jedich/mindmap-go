package models

// Map model
type Map struct {
	Model
	Name      string `json:"name" gorm:"default:Unnamed map"`
	Desc      string `json:"desc"`
	CreatorID int    `json:"creator_id"`
	Cards     []Card `json:"-" gorm:"foreignKey:MapID"`
}

type MapUpdate struct {
	Name        string
	Description string
}

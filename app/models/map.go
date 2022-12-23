package models

// Map model
type Map struct {
	Model
	Name      string `json:"name" query:"name" gorm:"default:Unnamed map"`
	Desc      string `json:"desc" query:"desc"`
	CreatorID int    `json:"creator_id" query:"creator_id"`
}

type MapUpdate struct {
	Name        string
	Description string
}

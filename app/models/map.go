package models

import "gorm.io/gorm"

// Map model
type Map struct {
	gorm.Model
	Name      string `json:"name,omitempty" query:"name" gorm:"default:Unnamed map"`
	Desc      string `json:"desc,omitempty" query:"desc"`
	CreatorID int    `json:"creator_id" query:"creator_id"`
}

type MapUpdate struct {
	Name        string
	Description string
}

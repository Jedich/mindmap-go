package models

// File model
type File struct {
	Model
	Filepath      string `json:"filename" gorm:"column:file_path"`
	FileExtension string `json:"extension" gorm:"column:file_extension"`
	Width         int    `json:"width"`
	Height        int    `json:"height"`
	CardID        int    `json:"card_id"`
}

type ImageForm struct {
	Filepath      string `json:"filepath" gorm:"column:file_path"`
	FileExtension string `json:"extension" gorm:"column:file_extension"`
	Width         int    `json:"width"`
	Height        int    `json:"height"`
}

package models

// File model
type File struct {
	Model
	Filepath      string `json:"filepath" gorm:"column:file_path"`
	FileExtension string `json:"extension" gorm:"column:file_extension"`
	Description   string `json:"description"`
	CardID        int    `json:"card_id"`
}

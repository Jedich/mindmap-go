package repository

import (
	"errors"
	"gorm.io/gorm"
	"mindmap-go/app/models"
	"mindmap-go/internal/database"
)

type CardRepo struct {
	DB *database.Database
}

type CardRepository interface {
	CreateCard(card *models.Card) error
	GetCardsByMapID(mapID int) ([]*models.Card, error)
	GetCardByID(id int) (*models.Card, error)
	UpdateCard(card *models.Card) error
	DeleteCard(card *models.Card) error
	CreateFile(file *models.File) error
}

func NewCardRepository(database *database.Database) CardRepository {
	return &CardRepo{
		DB: database,
	}
}

func (c *CardRepo) CreateCard(card *models.Card) error {
	return c.DB.Connection.Create(&card).Error
}

func (c *CardRepo) CreateFile(file *models.File) error {
	err := c.DB.Connection.Model(&models.File{}).Where("card_id = ?", file.CardID).First(&models.File{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.DB.Connection.Model(&models.File{}).Where("card_id = ?", file.CardID).Create(file).Error
		}
		return err
	}
	return c.DB.Connection.Model(&models.File{}).Where("card_id = ?", file.CardID).Updates(file).Error
}

func (c *CardRepo) GetCardsByMapID(mapID int) ([]*models.Card, error) {
	var res []*models.Card
	err := c.DB.Connection.Model(&models.Card{}).Preload("File").Where("map_id = ?", mapID).Find(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*models.Card{}, nil
		}
		return nil, err
	}
	return res, nil
}

func (c *CardRepo) GetCardByID(id int) (*models.Card, error) {
	res := &models.Card{}
	if err := c.DB.Connection.Where("id = ?", id).First(res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return res, nil
}

func (c *CardRepo) UpdateCard(card *models.Card) error {
	return c.DB.Connection.Omit("created_at", "deleted_at", "creator_id", "map_id").Save(&card).Error
}

func (c *CardRepo) DeleteCard(card *models.Card) error {
	return c.DB.Connection.Unscoped().Where("id = ? AND creator_id = ?", card.ID, card.CreatorID).Delete(&card).Error
}

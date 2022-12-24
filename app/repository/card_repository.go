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
	CreateCards(cards []*models.Card) error
	//GetAllByUser(userID int) ([]*models.Card, error)
	GetCardsByMapID(mapID int) ([]*models.Card, error)
	UpdateCard(card *models.Card) error
	DeleteCard(card *models.Card) error
}

func NewCardRepository(database *database.Database) CardRepository {
	return &CardRepo{
		DB: database,
	}
}

func (c *CardRepo) CreateCards(cards []*models.Card) error {
	return c.DB.Connection.Create(&cards).Error
}

func (c *CardRepo) GetCardsByMapID(mapID int) ([]*models.Card, error) {
	var res []*models.Card
	err := c.DB.Connection.Where("map_id = ?", mapID).Find(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*models.Card{}, nil
		}
		return nil, err
	}
	return res, nil
}

func (c *CardRepo) UpdateCard(card *models.Card) error {
	//TODO implement me
	panic("implement me")
}

func (c *CardRepo) DeleteCard(card *models.Card) error {
	//TODO implement me
	panic("implement me")
}

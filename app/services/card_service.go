package services

import (
	"mindmap-go/app/models"
	"mindmap-go/app/repository"
)

type CardSvc struct {
	Repo repository.CardRepository
}

type CardService interface {
	CreateCard(cardForm *CardForm) (*models.Card, error)
	GetCardsByMapID(mapID int) ([]*models.Card, error)
	GetCardByID(id int) (*models.Card, error)
	UpdateCard(card *models.CardUpdate) error
	DeleteCard(card *models.Card) error
}

func NewCardService(repo repository.CardRepository) CardService {
	return &CardSvc{
		Repo: repo,
	}
}

func (c *CardSvc) CreateCard(cardForm *CardForm) (*models.Card, error) {
	req := &models.Card{
		Name:      cardForm.Name,
		Text:      cardForm.Text,
		Level:     cardForm.Level,
		PositionY: cardForm.PositionY,
		ParentID:  cardForm.ParentID,
		CreatorID: cardForm.CreatorID,
		MapID:     cardForm.MapID,
	}
	err := c.Repo.CreateCard(req)
	return req, err
}

func (c *CardSvc) GetCardByID(id int) (*models.Card, error) {
	return c.Repo.GetCardByID(id)
}

func (c *CardSvc) GetCardsByMapID(mapID int) ([]*models.Card, error) {
	return c.Repo.GetCardsByMapID(mapID)
}

func (c *CardSvc) UpdateCard(card *models.CardUpdate) error {
	req, err := c.GetCardByID(card.ID)
	if err != nil {
		return err
	}

	req.Name = card.Name
	req.Text = card.Text
	req.Level = card.Level
	req.PositionY = card.PositionY
	req.ParentID = card.ParentID

	return c.Repo.UpdateCard(req)
}

func (c *CardSvc) DeleteCard(card *models.Card) error {
	return c.Repo.DeleteCard(card)
}

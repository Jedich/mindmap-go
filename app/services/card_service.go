package services

import (
	"mindmap-go/app/models"
	"mindmap-go/app/repository"
)

type CardSvc struct {
	Repo repository.CardRepository
}

type CardService interface {
	CreateCard(cardForm *CombinedCardForm) error
	GetCardsByMapID(mapID int) ([]*models.Card, error)
	UpdateCard(card *models.Card) error
	DeleteCard(card *models.Card) error
}

func NewCardService(repo repository.CardRepository) CardService {
	return &CardSvc{
		Repo: repo,
	}
}

func (c *CardSvc) CreateCard(cardForm *CombinedCardForm) error {
	req := make([]*models.Card, 0, len(cardForm.Cards))
	for _, card := range cardForm.Cards {
		req = append(req, &models.Card{
			Name:      card.Name,
			Text:      card.Text,
			ParentID:  card.ParentID,
			CreatorID: cardForm.CreatorID,
			MapID:     cardForm.MapID,
		})
	}
	return c.Repo.CreateCards(req)
}

func (c *CardSvc) GetCardsByMapID(mapID int) ([]*models.Card, error) {
	return c.Repo.GetCardsByMapID(mapID)
}

func (c *CardSvc) UpdateCard(card *models.Card) error {
	//TODO implement me
	panic("implement me")
}

func (c *CardSvc) DeleteCard(card *models.Card) error {
	//TODO implement me
	panic("implement me")
}

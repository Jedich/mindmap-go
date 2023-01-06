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
	GetCardsByMapID(mapID int) (*Component, error)
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
		Color:     cardForm.Color,
		ParentID:  cardForm.ParentID,
		CreatorID: cardForm.CreatorID,
		MapID:     cardForm.MapID,
		File:      cardForm.File,
	}
	err := c.Repo.CreateCard(req)
	return req, err
}

func (c *CardSvc) GetCardByID(id int) (*models.Card, error) {
	return c.Repo.GetCardByID(id)
}

func (c *CardSvc) GetCardsByMapID(mapID int) (*Component, error) {
	cards, err := c.Repo.GetCardsByMapID(mapID)
	if err != nil {
		return nil, err
	}

	cardMap := make(map[int]Component)
	//cardResp := make([]Component, 0, len(cards))
	var root Component

	for _, card := range cards {
		children := make([]Component, 0, 4)
		var res Component
		this := &CardTree{
			ID:       card.ID,
			Name:     card.Name,
			Text:     card.Text,
			Color:    card.Color,
			Children: children,
			ParentID: card.ParentID,
		}
		if card.File != nil {
			res = &CardWithFile{CardTree: *this, FIle: card.File}
		}
		if res == nil {
			res = this
		}
		cardMap[card.ID] = res
		//cardResp = append(cardResp, res)
	}
	for _, card := range cardMap {
		if card.getParentID() != nil {
			node := cardMap[*card.getParentID()]
			node.add(card)
		} else {
			root = card
		}
	}
	return &root, nil
}

func (c *CardSvc) UpdateCard(card *models.CardUpdate) error {
	req, err := c.GetCardByID(card.ID)
	if err != nil {
		return err
	}

	req.Name = card.Name
	req.Text = card.Text
	req.Color = card.Color
	req.ParentID = card.ParentID

	return c.Repo.UpdateCard(req)
}

func (c *CardSvc) DeleteCard(card *models.Card) error {
	return c.Repo.DeleteCard(card)
}

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
	GetCardsByMapID(mapID int) (*CardResponse, error)
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
	}
	err := c.Repo.CreateCard(req)
	return req, err
}

func (c *CardSvc) GetCardByID(id int) (*models.Card, error) {
	return c.Repo.GetCardByID(id)
}

func (c *CardSvc) GetCardsByMapID(mapID int) (*CardResponse, error) {
	cards, err := c.Repo.GetCardsByMapID(mapID)
	if err != nil {
		return nil, err
	}

	cardMap := make(map[int]*CardResponse)
	cardResp := make([]*CardResponse, 0, len(cards))
	var root *CardResponse

	for _, card := range cards {
		children := make([]*CardResponse, 0, 4)
		this := &CardResponse{
			Name:     card.Name,
			Text:     card.Text,
			Color:    card.Color,
			Children: children,
			ParentID: card.ParentID,
		}
		cardMap[card.ID] = this
		cardResp = append(cardResp, this)
	}
	for _, card := range cardResp {
		if card.ParentID != nil {
			node := cardMap[*card.ParentID]
			node.Children = append(node.Children, card)
		} else {
			root = card
		}
	}
	return root, nil
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

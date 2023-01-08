package services

import (
	"math/rand"
	"mindmap-go/app/models"
	"mindmap-go/app/repository"
	"mindmap-go/app/services/forms"
	"time"
)

type CardSvc struct {
	Repo repository.CardRepository
}

type CardService interface {
	CreateCard(cardForm *forms.CardForm) (*models.Card, error)
	GetCardsByMapID(mapID int) (*forms.Component, error)
	GetCardByID(id int) (*models.Card, error)
	UpdateCard(card *models.CardUpdate) error
	DeleteCard(card *models.Card) error
	GetRandomFilename(length int) string
}

func NewCardService(repo repository.CardRepository) CardService {
	return &CardSvc{
		Repo: repo,
	}
}

func (c *CardSvc) CreateCard(cardForm *forms.CardForm) (*models.Card, error) {
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

func (c *CardSvc) GetCardsByMapID(mapID int) (*forms.Component, error) {
	cards, err := c.Repo.GetCardsByMapID(mapID)
	if err != nil {
		return nil, err
	}

	cardMap := make(map[int]forms.Component)
	//cardResp := make([]Component, 0, len(cards))
	var root forms.Component

	for _, card := range cards {
		children := make([]forms.Component, 0, 4)
		var res forms.Component
		this := &forms.CardNode{
			ID:       card.ID,
			Name:     card.Name,
			Text:     card.Text,
			Color:    card.Color,
			Children: children,
			ParentID: card.ParentID,
		}
		if card.File != nil {
			res = &forms.CardNodeWithFile{CardNode: *this, FIle: card.File}
		}
		if res == nil {
			res = this
		}
		cardMap[card.ID] = res
		//cardResp = append(cardResp, res)
	}
	for _, card := range cardMap {
		if card.GetParentID() != nil {
			node := cardMap[*card.GetParentID()]
			node.Add(card)
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
	req.File = card.File

	return c.Repo.UpdateCard(req)
}

func (c *CardSvc) DeleteCard(card *models.Card) error {
	return c.Repo.DeleteCard(card)
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func (c *CardSvc) GetRandomFilename(length int) string {
	seededRand := rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

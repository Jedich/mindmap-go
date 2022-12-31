package controllers

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	"mindmap-go/app/models"
	"mindmap-go/app/services"
	"mindmap-go/utils/response"
	"strconv"
)

type Card struct {
	cardService services.CardService
}

type CardController interface {
	Index(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Destroy(c *fiber.Ctx) error
}

func NewCardController(cardService services.CardService) CardController {
	return &Card{
		cardService: cardService,
	}
}

func (card *Card) Index(c *fiber.Ctx) error {
	mapID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	res, err := card.cardService.GetCardsByMapID(mapID)
	if err != nil {
		return err
	}

	return response.Send(c, response.Body{
		Data: res,
	})
}

func (card *Card) Store(c *fiber.Ctx) error {
	tokenData := ParseToken(c)

	form := new(services.CardForm)

	if err := c.BodyParser(form); err != nil {
		return err
	}
	form.CreatorID = tokenData.id

	if err := validation.Validate(form); err != nil {
		return err
	}

	res, err := card.cardService.CreateCard(form)
	if err != nil {
		return err
	}

	return response.Send(c, response.Body{
		Messages: response.Messages{"Created!"},
		Data:     res,
	})
}

func (card *Card) Update(c *fiber.Ctx) error {

	form := new(models.CardUpdate)

	if err := c.BodyParser(form); err != nil {
		return err
	}

	if err := validation.Validate(form); err != nil {
		return err
	}

	if err := card.cardService.UpdateCard(form); err != nil {
		return err
	}

	return response.Send(c, response.Body{})
}

func (card *Card) Destroy(c *fiber.Ctx) error {
	mapID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	form := &models.Card{CreatorID: ParseToken(c).id, Model: models.Model{ID: mapID}}
	if err = card.cardService.DeleteCard(form); err != nil {
		return err
	}

	return response.Send(c, response.Body{})
}

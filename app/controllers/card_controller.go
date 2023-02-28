package controllers

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	"io"
	"mindmap-go/app/models"
	"mindmap-go/app/services"
	"mindmap-go/app/services/forms"
	"mindmap-go/utils/response"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Card struct {
	cardService services.CardService
}

type CardController interface {
	Index(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Destroy(c *fiber.Ctx) error

	StoreFile(c *fiber.Ctx) error
	UpdateFile(c *fiber.Ctx) error
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

	return response.NewResponseBuilder().
		WithData(res).Build().Send(c)
}

func (card *Card) Store(c *fiber.Ctx) error {
	tokenData := ParseToken(c)

	form := new(forms.CardForm)

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

	return response.NewResponseBuilder().
		WithMessages(response.Messages{"Created!"}).
		WithData(res).Build().Send(c)
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

	return response.NewResponseBuilder().Build().Send(c)
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

	return response.NewResponseBuilder().Build().Send(c)
}

func (card *Card) StoreFile(c *fiber.Ctx) error {
	cardID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	file, err := c.FormFile("uploadFile")
	if err != nil {
		return err
	}

	buffer, err := file.Open()
	if err != nil {
		return err
	}
	defer buffer.Close()
	fileSplit := strings.Split(file.Filename, ".")
	extension := fileSplit[len(fileSplit)-1]
	fname := card.cardService.GetRandomFilename(10)
	path := filepath.Join(".", "resources")
	_ = os.MkdirAll(path, os.ModePerm)
	fullPath := fmt.Sprintf("%s/%s.%s", path, fname, extension)

	f, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	defer func() {
		err = f.Close()
	}()
	// Copy the file to the destination path
	_, err = io.Copy(f, buffer)
	if err != nil {
		return err
	}

	if err = card.cardService.UpdateCardFile(&models.File{
		Filepath:      fmt.Sprintf("%s.%s", fname, extension),
		FileExtension: extension,
		CardID:        cardID,
	}); err != nil {
		return err
	}

	return response.NewResponseBuilder().WithData(map[string]interface{}{
		"filename": fmt.Sprintf("%s.%s", fname, extension),
	}).Build().Send(c)
}

func (card *Card) UpdateFile(c *fiber.Ctx) error {
	cardID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	form := new(models.ImageForm)

	if err := c.BodyParser(form); err != nil {
		return err
	}
	if err := card.cardService.UpdateCardFile(&models.File{
		Filepath:      form.Filepath,
		FileExtension: form.FileExtension,
		Width:         form.Width,
		Height:        form.Height,
		CardID:        cardID,
	}); err != nil {
		return err
	}

	return response.NewResponseBuilder().Build().Send(c)
}

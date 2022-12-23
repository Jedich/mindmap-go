package controllers

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	"mindmap-go/app/services"
	"mindmap-go/utils/response"
	"strconv"
)

type Map struct {
	mapService services.MapService
}

type MapController interface {
	Index(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Destroy(c *fiber.Ctx) error
}

func NewMapController(mapService services.MapService) MapController {
	return &Map{
		mapService: mapService,
	}
}

func (m *Map) Index(c *fiber.Ctx) error {
	tokenData := ParseToken(c)

	res, err := m.mapService.GetAllByUser(tokenData.id)
	if err != nil {
		return err
	}

	return response.Send(c, response.Body{
		Data: res,
	})
}

func (m *Map) Show(c *fiber.Ctx) error {
	tokenData := ParseToken(c)

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	res, err := m.mapService.GetMapByID(id, tokenData.id)
	if err != nil {
		return err
	}

	return response.Send(c, response.Body{
		Data: res,
	})
}

func (m *Map) Store(c *fiber.Ctx) error {
	tokenData := ParseToken(c)

	form := new(services.MapForm)

	if err := c.BodyParser(form); err != nil {
		return err
	}
	form.CreatorID = tokenData.id

	if err := validation.Validate(form); err != nil {
		return err
	}

	err := m.mapService.CreateMap(form)
	if err != nil {
		return err
	}

	return response.Send(c, response.Body{
		Messages: response.Messages{"Created!"},
	})
}

func (m *Map) Update(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (m *Map) Destroy(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

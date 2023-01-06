package controllers

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	"mindmap-go/app/models"
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

	mapsMap := make(map[string]*models.Map)
	for _, item := range res {
		mapsMap[strconv.Itoa(item.ID)] = item
	}

	return response.NewResponseBuilder().
		WithData(mapsMap).Build().Send(c)
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

	return response.NewResponseBuilder().
		WithData(res).Build().Send(c)
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

	res, err := m.mapService.CreateMap(form)
	if err != nil {
		return err
	}

	return response.NewResponseBuilder().
		WithMessages(response.Messages{"Created!"}).
		WithData(res).Build().Send(c)
}

func (m *Map) Update(c *fiber.Ctx) error {
	tokenData := ParseToken(c)

	form := new(models.MapUpdate)

	if err := c.BodyParser(form); err != nil {
		return err
	}

	if err := m.mapService.UpdateMap(form, tokenData.id); err != nil {
		return err
	}

	return response.NewResponseBuilder().Build().Send(c)
}

func (m *Map) Destroy(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

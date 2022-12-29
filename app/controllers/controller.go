package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"mindmap-go/app/services"
)

type Controller struct {
	User UserController
	Auth AuthController
}

func NewController(userService services.UserService, mapService services.MapService, cardService services.CardService) *Controller {
	return &Controller{
		User: &User{userService: userService},
		Auth: &Auth{userService: userService, mapService: mapService, cardService: cardService},
	}
}

func ParseToken(c *fiber.Ctx) *TokenData {
	data := c.Locals("user").(*jwt.Token)
	claims := data.Claims.(jwt.MapClaims)
	return &TokenData{int(claims["iss"].(float64))}
}

type TokenData struct {
	id int
}

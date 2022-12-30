package controllers

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"mindmap-go/app/models"
	"mindmap-go/app/services"
	"mindmap-go/utils"
	"mindmap-go/utils/response"
	"time"
)

type Auth struct {
	userService services.UserService
	mapService  services.MapService
	cardService services.CardService
}

type AuthController interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

func (a *Auth) Register(c *fiber.Ctx) error {

	registerForm := new(services.RegisterForm)

	if err := c.BodyParser(registerForm); err != nil {
		return err
	}

	if err := validation.Validate(registerForm); err != nil {
		return err
	}

	user, err := a.userService.Register(registerForm)
	if err != nil {
		return err
	}

	token, exp, err := CreateToken(user)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{Name: "token", Value: token, Expires: *exp, HTTPOnly: true})

	mindMap, err := a.mapService.CreateMap(&services.MapForm{CreatorID: user.ID})
	if err != nil {
		return err
	}

	cards, err := a.cardService.GetCardsByMapID(mindMap.ID)
	if err != nil {
		return err
	}

	return response.Send(c, response.Body{
		Messages: response.Messages{"The user was registered successfully!"},
		Data: map[string]interface{}{
			"user":   user,
			"newMap": mindMap,
			"tree":   cards,
		},
	})
}

func (a *Auth) Login(c *fiber.Ctx) error {

	cred := new(services.LoginForm)

	if err := c.BodyParser(cred); err != nil {
		return err
	}

	if err := validation.Validate(cred); err != nil {
		return err
	}

	user, err := a.userService.AuthorizeUser(cred)
	if err != nil {
		return err
	}

	token, exp, err := CreateToken(user)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{Name: "token", Value: token, Expires: *exp})

	mindMaps, err := a.mapService.GetAllByUser(user.ID)
	if err != nil {
		return err
	}

	var cards *services.CardResponse
	if len(mindMaps) > 0 {
		cards, err = a.cardService.GetCardsByMapID(mindMaps[0].ID)
		if err != nil {
			return err
		}
	}

	return response.Send(c, response.Body{
		Messages: response.Messages{"Logged in!"},
		Data: map[string]interface{}{
			"user":  user,
			"maps":  mindMaps,
			"cards": cards,
		},
	})
}

func CreateToken(user *models.User) (string, *time.Time, error) {
	exp := time.Now().Add(time.Hour * 72)
	// Create the Claims
	claims := jwt.MapClaims{
		"iss": user.ID,
		"exp": exp.Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(utils.ReadEnv("JWT_SECRET")))
	if err != nil {
		panic(err)
		return "", nil, err
	}

	return t, &exp, nil
}

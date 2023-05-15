package controllers

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"mindmap-go/app/models"
	"mindmap-go/app/services"
	"mindmap-go/app/services/forms"
	"mindmap-go/utils"
	"mindmap-go/utils/response"
	"strconv"
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

	registerForm := new(forms.RegisterForm)

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

	c.Cookie(&fiber.Cookie{Name: "token", Value: token, Expires: *exp})

	mindMap, err := a.mapService.CreateMap(&forms.MapForm{CreatorID: user.ID})
	if err != nil {
		return err
	}

	mindMaps := make(map[string]*models.Map)
	mindMaps[strconv.Itoa(mindMap.ID)] = mindMap

	return response.NewResponseBuilder().
		WithMessages(response.Messages{"The user was registered successfully!"}).
		WithData(map[string]interface{}{
			"user": user,
			"maps": mindMaps,
		}).Build().Send(c)
}

func (a *Auth) Login(c *fiber.Ctx) error {

	cred := new(forms.LoginForm)

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
	log.Println(user.ID)
	token, exp, err := CreateToken(user)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{Name: "token", Value: token, Expires: *exp})

	res, err := a.mapService.GetAllByUser(user.ID)
	if err != nil {
		return err
	}

	mindMaps := make(map[string]*models.Map)
	for _, item := range res {
		mindMaps[strconv.Itoa(item.ID)] = item
	}

	return response.NewResponseBuilder().
		WithMessages(response.Messages{"Logged in!"}).
		WithData(map[string]interface{}{
			"user": user,
			"maps": mindMaps,
		}).Build().Send(c)
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
		return "", nil, err
	}

	return t, &exp, nil
}

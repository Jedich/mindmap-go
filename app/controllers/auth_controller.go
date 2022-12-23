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

	token, err := CreateToken(user)
	if err != nil {
		return err
	}

	mindMap, err := a.mapService.CreateMap(&services.MapForm{CreatorID: int(user.ID)})
	if err != nil {
		return err
	}

	return response.Send(c, response.Body{
		Messages: response.Messages{"The user was registered successfully!"},
		Data: map[string]interface{}{
			"user":   user,
			"newMap": mindMap,
			"token":  token,
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

	token, err := CreateToken(user)
	if err != nil {
		return err
	}

	return response.Send(c, response.Body{
		Messages: response.Messages{"Logged in!"},
		Data: map[string]interface{}{
			"token": token,
		},
	})
}

func CreateToken(user *models.User) (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"iss": user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(utils.ReadEnv("JWT_SECRET")))
	if err != nil {
		panic(err)
		return "", err
	}

	return t, nil
}

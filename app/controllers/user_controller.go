package controllers

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"mindmap-go/app/models"
	"mindmap-go/app/services"
	"mindmap-go/utils/response"
	"strconv"
)

type User struct {
	userService services.UserService
}

type UserController interface {
	Index(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Destroy(c *fiber.Ctx) error
}

func NewUserController(userService services.UserService) UserController {
	return &User{
		userService: userService,
	}
}

func (u *User) Index(c *fiber.Ctx) error {

	res, err := u.userService.GetAllUsers()
	if err != nil {
		return err
	}

	return response.Send(c, response.Body{
		Data: res,
	})
}

func (u *User) Show(c *fiber.Ctx) error {

	result, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	res, err := u.userService.GetUserByID(result)
	if err != nil {
		return err
	}

	return response.Send(c, response.Body{
		Data: res,
	})
}

func (u *User) Update(c *fiber.Ctx) error {
	data := c.Locals("user").(*jwt.Token)
	claims := data.Claims.(jwt.MapClaims)
	id := claims["iss"].(int)

	user, err := u.userService.GetUserByID(id)
	if err != nil {
		return err
	}

	req := new(models.UserUpdate)

	if err = c.BodyParser(user); err != nil {
		return err
	}

	if err = validation.Validate(user); err != nil {
		return err
	}

	if err = u.userService.UpdateUser(user, req); err != nil {
		return err
	}

	return response.Send(c, response.Body{
		Messages: response.Messages{"The user was updated successfully!"},
	})
}

func (u *User) Destroy(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

package controllers

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
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

	return response.NewResponseBuilder().
		WithData(res).Build().Send(c)
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

	return response.NewResponseBuilder().
		WithData(res).Build().Send(c)
}

func (u *User) Update(c *fiber.Ctx) error {
	tokenData := ParseToken(c)

	user, err := u.userService.GetUserByID(tokenData.id)
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

	return response.NewResponseBuilder().
		WithMessages(response.Messages{"The user was updated successfully!"}).
		Build().Send(c)
}

func (u *User) Destroy(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

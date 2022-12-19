package controllers

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"mindmap-go/app/services"
	"mindmap-go/utils"
	"mindmap-go/utils/response"
)

type User struct {
	userService services.UserService
}

type UserController interface {
	Index(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Destroy(c *fiber.Ctx) error
}

func NewUserController(userService services.UserService) UserController {
	return &User{
		userService: userService,
	}
}

func (u *User) Index(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (u *User) Show(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (u *User) Store(c *fiber.Ctx) error {

	user := new(services.UserForm)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	if err := validation.Validate(user); err != nil {
		return err
	}

	res, err := u.userService.Register(user)
	if err != nil {
		if errors.Is(err, utils.DuplicateEntryError) {
			return response.Send(c, response.Body{
				Code:     fiber.StatusBadRequest,
				Messages: response.Messages{"The user with such credentials already exists."},
			})
		}
		return err
	}

	return response.Send(c, response.Body{
		Messages: response.Messages{"The user was registered successfully!"},
		Data:     res,
	})
}

func (u *User) Update(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (u *User) Destroy(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

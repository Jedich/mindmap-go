package controllers

import (
	"mindmap-go/app/services"
)

type Controller struct {
	User UserController
	Auth AuthController
}

func NewController(userService services.UserService) *Controller {
	return &Controller{
		User: &User{userService: userService},
		Auth: &Auth{userService: userService},
	}
}

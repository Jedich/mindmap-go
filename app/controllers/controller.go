package controllers

import (
	"mindmap-go/app/services"
)

type Controller struct {
	User UserController
}

func NewController(userService services.UserService) *Controller {
	return &Controller{
		User: &User{userService: userService},
	}
}

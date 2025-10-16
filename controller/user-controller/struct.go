package user_controller

import (
	"backendmailingroom/controller"
	"backendmailingroom/repository"
)

type UserHandler struct {
	user repository.UserRepository
}

func NewUserController(user repository.UserRepository) controller.UserController {
	return &UserHandler{
		user: user,
	}
}

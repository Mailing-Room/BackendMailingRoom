package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	InputUser(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	GetAllUsers(c *fiber.Ctx) error
	GetUserByID(c *fiber.Ctx) error
	GetUserByEmail(c *fiber.Ctx) error
	DeleteUserByID(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
}

type DepartemenController interface {
	InputDepartemen(c *fiber.Ctx) error
}

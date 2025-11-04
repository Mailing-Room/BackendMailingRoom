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

type SubdirektoratController interface {
	InputSubdirektorat(c *fiber.Ctx) error
}

type OfficeController interface {
	InputOffice(c *fiber.Ctx) error
	GetOfficeByID(c *fiber.Ctx) error
	GetAllOffice(c *fiber.Ctx) error
	GetOfficeByKota(c *fiber.Ctx) error
	DeleteOfficeByID(c *fiber.Ctx) error
	UpdateOffice(c *fiber.Ctx) error
}

type NaskahController interface {
}

type KategoriController interface {
}

type DivisiController interface {
}

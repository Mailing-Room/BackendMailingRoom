package office_controller

import (
	"backendmailingroom/model"

	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (o *OfficeHandler) InputOffice(c *fiber.Ctx) error {
	var office model.Office
	if err := c.BodyParser(&office); err != nil {
		return fmt.Errorf("failed to parse request body: %w", err)
	}

	if office.NamaOffice == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Nama office tidak boleh kosong",
		})
	}

	if office.Alamat == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Alamat tidak boleh kosong",
		})
	}

	if office.KodePos == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Kode pos office tidak boleh kosong",
		})
	}

	if !validateKodePos(office.KodePos) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Kode pos office tidak valid",
		})
	}

	result, err := o.office.InputOffice(c.Context(), office)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Gagal menambahkan office: %v", err),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Office berhasil ditambahkan",
		"office":  result,
	})
}

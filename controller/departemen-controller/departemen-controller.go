package departemen_controller

import (
	"backendmailingroom/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (d *DepartemenHandler) InputDepartemen(c *fiber.Ctx) error {
	var departemen model.Departemen
	if err := c.BodyParser(&departemen); err != nil {
		return fmt.Errorf("failed to parse request body: %w", err)
	}

	// Validasi input wajib
	if departemen.NamaDepartemen == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Nama departemen tidak boleh kosong",
		})
	}

	if departemen.KodeDepartemen == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Kode departemen tidak boleh kosong",
		})
	}

	if departemen.NoTelp == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Nomor telepon departemen tidak boleh kosong",
		})
	}

	// Simpan ke database via repository
	result, err := d.departemen.InputDepartemen(c.Context(), departemen)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Gagal menambahkan departemen: %v", err),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":     "success",
		"message":    "Departemen berhasil ditambahkan",
		"departemen": result,
	})
}

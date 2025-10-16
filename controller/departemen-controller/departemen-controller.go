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

	// Validasi input
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

	if departemen.Alamat == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Alamat departemen tidak boleh kosong",
		})
	}

	if departemen.NoTelp == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "No. Telp departemen tidak boleh kosong",
		})
	}

	// Validasi kode pos
	if departemen.KodePos == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Kode pos departemen tidak boleh kosong",
		})
	}

	if !validateKodePos(departemen.KodePos) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Kode pos tidak valid. Kode pos harus berupa angka dengan panjang 5 digit.",
		})
	}

	result, err := d.departemen.InputDepartemen(c.Context(), departemen)
	if err != nil {
		return fmt.Errorf("failed to create departemen: %w", err)
	}

	return c.JSON(fiber.Map{
		"status":     "success",
		"message":    "Departemen berhasil ditambahkan",
		"departemen": result,
	})
}

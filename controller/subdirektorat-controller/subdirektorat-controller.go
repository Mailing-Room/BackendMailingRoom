package departemen_controller

import (
	"backendmailingroom/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (d *SubdirektoratHandler) InputSubdirektorat(c *fiber.Ctx) error {
	var subdirektorat model.SubDirektorat
	if err := c.BodyParser(&subdirektorat); err != nil {
		return fmt.Errorf("failed to parse request body: %w", err)
	}

	// Validasi input wajib
	if subdirektorat.NamaSubDirektorat == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Nama subdirektorat tidak boleh kosong",
		})
	}

	if subdirektorat.KodeSubDirektorat == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Kode subdirektorat tidak boleh kosong",
		})
	}

	if subdirektorat.NoTelp == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Nomor telepon subdirektorat tidak boleh kosong",
		})
	}

	// Simpan ke database via repository
	result, err := d.subdirektorat.InputSubDirektorat(c.Context(), subdirektorat)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Gagal menambahkan subdirektorat: %v", err),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":        "success",
		"message":       "Subdirektorat berhasil ditambahkan",
		"subdirektorat": result,
	})
}

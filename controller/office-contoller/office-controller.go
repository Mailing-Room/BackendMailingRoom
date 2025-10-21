package office_controller

import (
	"backendmailingroom/model"
	"log"
	"net/http"
	"strings"

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

func (o *OfficeHandler) GetOfficeByID(c *fiber.Ctx) error {
	id := c.Params("id")
	office, err := o.office.GetOfficeByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Gagal mendapatkan office: %v", err),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"office": office,
	})
}

func (o *OfficeHandler) GetAllOffice(c *fiber.Ctx) error {
	offices, err := o.office.GetAllOffice(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Gagal mendapatkan data office: %v", err),
		})
	}

	// Kalau tidak ada data office
	if len(offices) == 0 {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Tidak ada data office tersedia",
			"offices": []model.Office{},
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Data office berhasil diambil",
		"offices": offices,
	})
}

func (o *OfficeHandler) GetOfficeByKota(c *fiber.Ctx) error {
	kota := c.Params("kota")
	if kota == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "kota parameter is required",
		})
	}

	offices, err := o.office.GetOfficeByKota(c.Context(), kota)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Gagal mendapatkan office: %v", err),
		})
	}

	if len(offices) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Office dengan kota " + kota + " tidak ditemukan",
			"data":    []model.Office{},
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Data office berhasil diambil",
		"data":    offices,
	})
}

func (o *OfficeHandler) DeleteOfficeByID(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "id parameter is required",
		})
	}

	office, err := o.office.DeleteOfficeByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Gagal menghapus office: %v", err),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Office berhasil dihapus",
		"office":  office,
	})
}

func (o *OfficeHandler) UpdateOffice(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "id parameter is required",
		})
	}

	var updatedData model.Office
	if err := c.BodyParser(&updatedData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal memparsing data: " + err.Error(),
		})
	}

	updatedOffice, err := o.office.UpdateOffice(c.Context(), id, updatedData)
	if err != nil {

		log.Printf("[ERROR] gagal update office (id: %s): %v", id, err)

		if strings.Contains(err.Error(), "tidak ditemukan") {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	log.Printf("[INFO] Office dengan id %s berhasil diperbarui", id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Office berhasil diperbarui",
		"office":  updatedOffice,
	})
}

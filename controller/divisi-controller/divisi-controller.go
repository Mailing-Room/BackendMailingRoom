package divisi_controller

import (
	"backendmailingroom/model"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (d *DivisiHandler) InputDivisi(c *fiber.Ctx) error {
	var divisi model.Divisi

	if err := c.BodyParser(&divisi); err != nil {
		log.Println("[ERROR] Gagal membaca body request:", err)
		return fmt.Errorf("failed to parse request body: %w", err)
	}

	// Validasi input wajib
	// if divisi.SubDirektoratID == "" {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"status":  "error",
	// 		"message": "SubDirektoratID tidak boleh kosong",
	// 	})
	// }

	if divisi.NamaDivisi == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Nama divisi tidak boleh kosong",
		})
	}

	if divisi.KodeDivisi == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Kode divisi tidak boleh kosong",
		})
	}

	result, err := d.divisi.InputDivisi(c.Context(), divisi)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Gagal menambahkan divisi: %v", err),
		})
	}

	log.Println("[INFO] Divisi berhasil ditambahkan dengan ID:", result.DivisiID)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Divisi berhasil ditambahkan",
		"divisi":  result,
	})
}

func (d *DivisiHandler) GetDivisiByID(c *fiber.Ctx) error {
	id := c.Params("id")

	divisi, err := d.divisi.GetDivisiByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Divisi tidak ditemukan: %v", err),
		})
	}

	log.Println("[INFO] Berhasil mengambil data divisi dengan ID:", id)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Berhasil mengambil data divisi",
		"divisi":  divisi,
	})
}

func (d *DivisiHandler) GetAllDivisi(c *fiber.Ctx) error {
	divisions, err := d.divisi.GetAllDivisi(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Gagal mengambil data divisi: %v", err),
		})
	}

	log.Println("[INFO] Berhasil mengambil semua data divisi, total:", len(divisions))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Berhasil mengambil semua data divisi",
		"data":    divisions,
		"total":   len(divisions),
	})
}

func (d *DivisiHandler) GetDivisiBySubDirektoratID(c *fiber.Ctx) error {
	subDirektoratID := c.Params("sub_direktorat_id")
	if subDirektoratID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "SubDirektorat ID parameter is required",
		})
	}

	divisions, err := d.divisi.GetDivisiBySubDirektoratID(c.Context(), subDirektoratID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Gagal mengambil data divisi: %v", err),
		})
	}

	if divisions == nil {
		divisions = []model.Divisi{}
	}

	log.Printf("[INFO] Berhasil mengambil data divisi untuk SubDirektorat ID %s, total: %d", subDirektoratID, len(divisions))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Berhasil mengambil data divisi",
		"data":    divisions,
		"total":   len(divisions),
	})
}

func (d *DivisiHandler) GetDivisiBySubDirektoratName(c *fiber.Ctx) error {
	namaSubDirektorat := c.Query("nama")
	if namaSubDirektorat == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Nama SubDirektorat query parameter is required",
		})
	}

	divisions, err := d.divisi.GetDivisiBySubDirektoratName(c.Context(), namaSubDirektorat)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Gagal mengambil data divisi: %v", err),
		})
	}

	// Jika tidak ada data, kembalikan array kosong
	if divisions == nil {
		divisions = []model.Divisi{}
	}

	log.Printf("[INFO] Ditemukan %d divisi untuk SubDirektorat: %s", len(divisions), namaSubDirektorat)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": fmt.Sprintf("Data divisi untuk Sub Direktorat '%s' berhasil diambil", namaSubDirektorat),
		"data":    divisions,
		"total":   len(divisions),
	})
}

func (d *DivisiHandler) DeleteDivisiByID(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "id parameter is required",
		})
	}

	divisi, err := d.divisi.DeleteDivisiByID(c.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), "tidak ditemukan") {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("failed to delete divisi: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "divisi deleted successfully",
		"divisi":  divisi,
	})
}

func (d *DivisiHandler) UpdateDivisi(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "id parameter is required",
		})
	}

	var divisi model.Divisi
	if err := c.BodyParser(&divisi); err != nil {
		return fmt.Errorf("failed to parse request body: %w", err)
	}

	// Validasi input wajib
	// if divisi.SubDirektoratID == "" {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"status":  "error",
	// 		"message": "Sub Direktorat ID tidak boleh kosong",
	// 	})
	// }

	if divisi.NamaDivisi == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Nama divisi tidak boleh kosong",
		})
	}

	if divisi.KodeDivisi == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Kode divisi tidak boleh kosong",
		})
	}

	// Update data via repository
	result, err := d.divisi.UpdateDivisi(c.Context(), id, divisi)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Gagal memperbarui divisi: %v", err),
		})
	}

	log.Printf("[INFO] Divisi dengan ID %s berhasil diperbarui", id)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Divisi berhasil diperbarui",
		"divisi":  result,
	})
}

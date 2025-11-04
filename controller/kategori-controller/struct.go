package kategori_controller

import (
	"backendmailingroom/controller"
	"backendmailingroom/repository"
)

type KategoriHandler struct {
	kategori repository.KategoriRepository
}

func NewKategoriController(kategori repository.KategoriRepository) controller.KategoriController {
	return &KategoriHandler{
		kategori: kategori,
	}
}

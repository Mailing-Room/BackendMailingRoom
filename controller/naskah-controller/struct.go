package naskah_controller

import (
	"backendmailingroom/controller"
	"backendmailingroom/repository"
)

type NaskahHandler struct {
	naskah repository.NaskahRepository
}

func NewNaskahController(naskah repository.NaskahRepository) controller.NaskahController {
	return &NaskahHandler{
		naskah: naskah,
	}
}

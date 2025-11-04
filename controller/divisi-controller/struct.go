package divisi_controller

import (
	"backendmailingroom/controller"
	"backendmailingroom/repository"
)

type DivisiHandler struct {
	divisi repository.DivisiRepository
}

func NewDivisiController(divisi repository.DivisiRepository) controller.DivisiController {
	return &DivisiHandler{
		divisi: divisi,
	}
}

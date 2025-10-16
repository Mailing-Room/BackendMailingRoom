package departemen_controller

import (
	"backendmailingroom/controller"
	"backendmailingroom/repository"
)

type DepartemenHandler struct {
	departemen repository.DepartemenRepository
}

func NewDepartemenController(departemen repository.DepartemenRepository) controller.DepartemenController {
	return &DepartemenHandler{
		departemen: departemen,
	}
}

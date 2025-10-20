package office_controller

import (
	"backendmailingroom/controller"
	"backendmailingroom/repository"
)

type OfficeHandler struct {
	office repository.OfficeRepository
}

func NewOfficeController(office repository.OfficeRepository) controller.OfficeController {
	return &OfficeHandler{
		office: office,
	}
}

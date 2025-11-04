package departemen_controller

import (
	"backendmailingroom/controller"
	"backendmailingroom/repository"
)

type SubdirektoratHandler struct {
	subdirektorat repository.SubdirektoratRepository
}

func NewSubdirektoratController(subdirektorat repository.SubdirektoratRepository) controller.SubdirektoratController {
	return &SubdirektoratHandler{
		subdirektorat: subdirektorat,
	}
}

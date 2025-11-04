package routes

import "backendmailingroom/repository"

// Declare repository variables
var (
	UserRepository          repository.UserRepository
	SubdirektoratRepository repository.SubdirektoratRepository
	OfficeRepository        repository.OfficeRepository
	NaskahRepository        repository.NaskahRepository
	KategoriRepository      repository.KategoriRepository
	DivisiRepository        repository.DivisiRepository
)

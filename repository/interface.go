package repository

import (
	"backendmailingroom/model"
	"context"
)

type UserRepository interface {
	InputUser(ctx context.Context, user model.User) (model.User, error)
	GetUserForLogin(ctx context.Context, email string) (model.User, error)
	GetAllUsers(ctx context.Context) ([]model.User, error)
	GetUserByID(ctx context.Context, id string) (model.User, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	DeleteUserByID(ctx context.Context, id string) (model.User, error)
	UpdateUser(ctx context.Context, id string, updatedData model.User) (model.User, error)
}

type SubdirektoratRepository interface {
	InputSubDirektorat(ctx context.Context, subdirektorat model.SubDirektorat) (model.SubDirektorat, error)
}

type OfficeRepository interface {
	InputOffice(ctx context.Context, office model.Office) (model.Office, error)
	GetOfficeByID(ctx context.Context, id string) (model.Office, error)
	GetAllOffice(ctx context.Context) ([]model.Office, error)
	GetOfficeByKota(ctx context.Context, kota string) ([]model.Office, error)
	DeleteOfficeByID(ctx context.Context, id string) (model.Office, error)
	UpdateOffice(ctx context.Context, id string, updatedData model.Office) (model.Office, error)
}

type NaskahRepository interface {
}

type KategoriRepository interface {
}

type DivisiRepository interface {
}

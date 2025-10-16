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

type DepartemenRepository interface {
	InputDepartemen(ctx context.Context, departemen model.Departemen) (model.Departemen, error)
}

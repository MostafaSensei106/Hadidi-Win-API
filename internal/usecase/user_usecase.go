package usecase

import (
	"context"

	"github.com/MostafaSensei106/Hadidi-Win-API/internal/domain"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

// CreateUser implements [domain.UserUsecase].
func (u *userUsecase) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	panic("unimplemented")
}

// DeleteUser implements [domain.UserUsecase].
func (u *userUsecase) DeleteUser(ctx context.Context, id string) error {
	panic("unimplemented")
}

// GetUser implements [domain.UserUsecase].
func (u *userUsecase) GetUser(ctx context.Context, id string) (*domain.User, error) {
	panic("unimplemented")
}

// UpdateUser implements [domain.UserUsecase].
func (u *userUsecase) UpdateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	panic("unimplemented")
}

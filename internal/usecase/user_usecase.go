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

func (u *userUsecase) GetUsers(ctx context.Context) ([]*domain.User, error) {
	return u.userRepo.GetAll(ctx)
}

func (u *userUsecase) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	return u.userRepo.Create(ctx, user)
}

func (u *userUsecase) DeleteUser(ctx context.Context, id string) error {
	return u.userRepo.Delete(ctx, id)
}

func (u *userUsecase) GetUser(ctx context.Context, id string) (*domain.User, error) {
	return u.userRepo.Get(ctx, id)
}

func (u *userUsecase) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	return u.userRepo.Update(ctx, user)
}

package repository

import (
	"context"

	"github.com/MostafaSensei106/Hadidi-Win-API/internal/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	BaseRepository[domain.User]
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		BaseRepository: NewBaseRepository[domain.User](db),
	}
}

func (r *userRepository) Get(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	if err := r.GetDB().WithContext(ctx).Preload("Addresses").First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetAll(ctx context.Context) ([]*domain.User, error) {
	var users []*domain.User
	if err := r.GetDB().WithContext(ctx).Preload("Addresses").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	return r.BaseRepository.Create(ctx, user)
}

func (r *userRepository) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	return r.BaseRepository.Update(ctx, user)
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	return r.BaseRepository.Delete(ctx, id)
}

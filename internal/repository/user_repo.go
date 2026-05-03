package repository

import (
	"context"

	"github.com/MostafaSensei106/Hadidi-Win-API/internal/domain"
	"github.com/redis/go-redis/v9"
)

type userRepository struct {
	db    *db.query
	redis *redis.Client
}

func NewUserRepository(db *db.query, radis *redis.Client) domain.UserRepository {
	return &userRepository{
		db:    db,
		redis: radis,
	}
}

// Create implements [domain.UserRepository].
func (u *userRepository) Create(ctx context.Context, user domain.User) (*domain.User, error) {
	panic("unimplemented")
}

// Delete implements [domain.UserRepository].
func (u *userRepository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// Get implements [domain.UserRepository].
func (u *userRepository) Get(ctx context.Context, id string) (*domain.User, error) {
	panic("unimplemented")
}

// Update implements [domain.UserRepository].
func (u *userRepository) Update(ctx context.Context, user domain.User) (*domain.User, error) {
	panic("unimplemented")
}

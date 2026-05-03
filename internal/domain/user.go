package domain

import (
	"context"
)

type User struct {
	Person
	AvatarURL string            `json:"avatar_url" gorm:"size:255"`
	Wishlist  []string          `json:"wishlist" gorm:"type:text[]"`
	Addresses []ShippingAddress `json:"addresses" gorm:"foreignKey:UserID"`
}

type UserRepository interface {
	Get(ctx context.Context, id string) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)
	Create(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	Delete(ctx context.Context, id string) error
}

type UserUsecase interface {
	GetUser(ctx context.Context, id string) (*User, error)
	GetUsers(ctx context.Context) ([]*User, error)
	CreateUser(ctx context.Context, user *User) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
	DeleteUser(ctx context.Context, id string) error
}

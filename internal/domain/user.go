package domain

import "context"

type User struct {
	AvatarURL string   `json:"avatar_url"`
	Wishlist  []string `json:"wishlist"`
	Person
	Addresses ShippingAddresses
}

type UserRepository interface {
	Get(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, user User) (*User, error)
	Update(ctx context.Context, user User) (*User, error)
	Delete(ctx context.Context, id string) error
}

type UserUsecase interface {
	GetUser(ctx context.Context, id string) (*User, error)
	CreateUser(ctx context.Context, user User) (*User, error)
	UpdateUser(ctx context.Context, user User) (*User, error)
	DeleteUser(ctx context.Context, id string) error
}

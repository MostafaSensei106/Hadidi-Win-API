package domain

import "context"

type Product struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Desc   string  `json:"desc"`
	Width  float32 `json:"width"`
	Height float32 `json:"height"`
	Color  string  `json:"color"`
	Price  int64   `json:"price_minor_units"`
}

type Window struct {
	Product
	WindowType string `json:"windowType"`
	GlassType  string `json:"glassType"`
	Brand      string `json:"brand"`
}

type Door struct {
	Product
	DoorType string `json:"doorType"`
	Brand    string `json:"brand"`
}

type ProductRepository interface {
	GetByID(ctx context.Context, id string) (*Product, error)
	GetAll(ctx context.Context) ([]*Product, error)
	Create(ctx context.Context, product *Product) (*Product, error)
	Update(ctx context.Context, product *Product) (*Product, error)
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, query string) ([]*Product, error)
}

type ProductUsecase interface {
	GetProduct(ctx context.Context, id string) (*Product, error)
	GetProducts(ctx context.Context) ([]*Product, error)
	CreateProduct(ctx context.Context, product *Product) (*Product, error)
	UpdateProduct(ctx context.Context, product *Product) (*Product, error)
	DeleteProduct(ctx context.Context, id string) error
	SearchProducts(ctx context.Context, query string) ([]*Product, error)
}

package repository

import (
	"context"

	"github.com/MostafaSensei106/Hadidi-Win-API/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

type productRepository struct {
	db    *db.query
	redis *redis.Client
}

func NewProductRepository(db *pgx.Conn, redis *redis.Client) domain.ProductRepository {
	return &productRepository{
		db:    db,
		redis: redis,
	}
}

// Create implements [domain.ProductRepository].
func (p *productRepository) Create(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	panic("unimplemented")
}

// Delete implements [domain.ProductRepository].
func (p *productRepository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// GetAll implements [domain.ProductRepository].
func (p *productRepository) GetAll(ctx context.Context) ([]*domain.Product, error) {
	panic("unimplemented")
}

// GetByID implements [domain.ProductRepository].
func (p *productRepository) GetByID(ctx context.Context, id string) (*domain.Product, error) {
	panic("unimplemented")
}

// Search implements [domain.ProductRepository].
func (p *productRepository) Search(ctx context.Context, query string) ([]*domain.Product, error) {
	panic("unimplemented")
}

// Update implements [domain.ProductRepository].
func (p *productRepository) Update(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	panic("unimplemented")
}

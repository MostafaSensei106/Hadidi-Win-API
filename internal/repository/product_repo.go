package repository

import (
	"context"

	"github.com/MostafaSensei106/Hadidi-Win-API/internal/domain"
	"gorm.io/gorm"
)

type productRepository struct {
	BaseRepository[domain.Product]
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &productRepository{
		BaseRepository: NewBaseRepository[domain.Product](db),
	}
}

func (r *productRepository) Get(ctx context.Context, id string) (*domain.Product, error) {
	return r.BaseRepository.Get(ctx, id)
}

func (r *productRepository) GetAll(ctx context.Context) ([]*domain.Product, error) {
	return r.BaseRepository.GetAll(ctx)
}

func (r *productRepository) Create(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	return r.BaseRepository.Create(ctx, product)
}

func (r *productRepository) Update(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	return r.BaseRepository.Update(ctx, product)
}

func (r *productRepository) Delete(ctx context.Context, id string) error {
	return r.BaseRepository.Delete(ctx, id)
}

func (r *productRepository) Search(ctx context.Context, query string) ([]*domain.Product, error) {
	var products []*domain.Product
	if err := r.GetDB().WithContext(ctx).Where("name ILIKE ? OR desc ILIKE ?", "%"+query+"%", "%"+query+"%").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

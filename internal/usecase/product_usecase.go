package usecase

import (
	"context"

	"github.com/MostafaSensei106/Hadidi-Win-API/internal/domain"
)

type productUsecase struct {
	productRepo domain.ProductRepository
}

func NewProductUsecase(productRepo domain.ProductRepository) domain.ProductUsecase {
	return &productUsecase{
		productRepo: productRepo,
	}
}

func (u *productUsecase) GetProduct(ctx context.Context, id string) (*domain.Product, error) {
	return u.productRepo.Get(ctx, id)
}

func (u *productUsecase) GetProducts(ctx context.Context) ([]*domain.Product, error) {
	return u.productRepo.GetAll(ctx)
}

func (u *productUsecase) CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	return u.productRepo.Create(ctx, product)
}

func (u *productUsecase) UpdateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	return u.productRepo.Update(ctx, product)
}

func (u *productUsecase) DeleteProduct(ctx context.Context, id string) error {
	return u.productRepo.Delete(ctx, id)
}

func (u *productUsecase) SearchProducts(ctx context.Context, query string) ([]*domain.Product, error) {
	return u.productRepo.Search(ctx, query)
}

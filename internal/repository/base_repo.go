package repository

import (
	"context"

	"gorm.io/gorm"
)

type BaseRepository[T any] interface {
	Create(ctx context.Context, entity *T) (*T, error)
	Get(ctx context.Context, id string) (*T, error)
	GetAll(ctx context.Context) ([]*T, error)
	Update(ctx context.Context, entity *T) (*T, error)
	Delete(ctx context.Context, id string) error
	GetDB() *gorm.DB
}

type baseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{
		db: db,
	}
}

func (r *baseRepository[T]) GetDB() *gorm.DB {
	return r.db
}

func (r *baseRepository[T]) Create(ctx context.Context, entity *T) (*T, error) {
	if err := r.db.WithContext(ctx).Create(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *baseRepository[T]) Delete(ctx context.Context, id string) error {
	var entity T
	return r.db.WithContext(ctx).Delete(&entity, "id = ?", id).Error
}

func (r *baseRepository[T]) Get(ctx context.Context, id string) (*T, error) {
	var entity T
	if err := r.db.WithContext(ctx).First(&entity, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *baseRepository[T]) GetAll(ctx context.Context) ([]*T, error) {
	var entities []*T
	if err := r.db.WithContext(ctx).Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *baseRepository[T]) Update(ctx context.Context, entity *T) (*T, error) {
	if err := r.db.WithContext(ctx).Save(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

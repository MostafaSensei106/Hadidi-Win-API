package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/MostafaSensei106/Hadidi-Win-API/internal/domain"
	"github.com/redis/go-redis/v9"
)

type cachedUserRepository struct {
	repo  domain.UserRepository
	redis *redis.Client
	ttl   time.Duration
}

func NewCachedUserRepository(repo domain.UserRepository, redis *redis.Client, ttl time.Duration) domain.UserRepository {
	return &cachedUserRepository{
		repo:  repo,
		redis: redis,
		ttl:   ttl,
	}
}

func (c *cachedUserRepository) Get(ctx context.Context, id string) (*domain.User, error) {
	key := fmt.Sprintf("user:%s", id)
	
	val, err := c.redis.Get(ctx, key).Result()
	if err == nil {
		var user domain.User
		if err := json.Unmarshal([]byte(val), &user); err == nil {
			return &user, nil
		}
	}

	user, err := c.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	data, _ := json.Marshal(user)
	c.redis.Set(ctx, key, data, c.ttl)

	return user, nil
}

func (c *cachedUserRepository) GetAll(ctx context.Context) ([]*domain.User, error) {
	return c.repo.GetAll(ctx)
}

func (c *cachedUserRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	u, err := c.repo.Create(ctx, user)
	if err == nil {
		key := fmt.Sprintf("user:%s", u.ID)
		data, _ := json.Marshal(u)
		c.redis.Set(ctx, key, data, c.ttl)
	}
	return u, err
}

func (c *cachedUserRepository) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	u, err := c.repo.Update(ctx, user)
	if err == nil {
		key := fmt.Sprintf("user:%s", u.ID)
		data, _ := json.Marshal(u)
		c.redis.Set(ctx, key, data, c.ttl)
	}
	return u, err
}

func (c *cachedUserRepository) Delete(ctx context.Context, id string) error {
	err := c.repo.Delete(ctx, id)
	if err == nil {
		key := fmt.Sprintf("user:%s", id)
		c.redis.Del(ctx, key)
	}
	return err
}

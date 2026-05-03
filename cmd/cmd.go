package cmd

import (
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/MostafaSensei106/Hadidi-Win-API/internal/delivery/https"
	"github.com/MostafaSensei106/Hadidi-Win-API/internal/env"
	"github.com/MostafaSensei106/Hadidi-Win-API/internal/middleware"
	"github.com/MostafaSensei106/Hadidi-Win-API/internal/repository"
	"github.com/MostafaSensei106/Hadidi-Win-API/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Execute() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg := config{
		port: ":8080",
		db: dbConfig{
			dns: env.GetString("Hadidi-Win-DB-DSN", "host=localhost user=root password=root dbname=hadidiwin sslmode=disable"),
		},
	}

	// Database (GORM)
	db, err := gorm.Open(postgres.Open(cfg.db.dns), &gorm.Config{})
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		panic(err)
	}

	redisClint := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Repositories
	productRepo := repository.NewProductRepository(db)
	userRepo := repository.NewUserRepository(db)

	// Caching Layer (Decorator)
	cachedUserRepo := repository.NewCachedUserRepository(userRepo, redisClint, 30*time.Minute)

	// Usecases
	productUsecase := usecase.NewProductUsecase(productRepo)
	userUsecase := usecase.NewUserUsecase(cachedUserRepo)

	r := gin.Default()
	r.Use(middleware.RateLimiter())

	https.NewProductHandler(r, productUsecase)
	https.NewUserHandler(r, userUsecase)

	log.Printf("Server is Starting at %s", cfg.port)
	r.Run(cfg.port)
}

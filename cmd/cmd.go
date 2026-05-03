package cmd

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/MostafaSensei106/Hadidi-Win-API/internal/delivery/https"
	"github.com/MostafaSensei106/Hadidi-Win-API/internal/env"
	"github.com/MostafaSensei106/Hadidi-Win-API/internal/middleware"
	"github.com/MostafaSensei106/Hadidi-Win-API/internal/repository"
	"github.com/MostafaSensei106/Hadidi-Win-API/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

func Execute() {
	ctx := context.Background()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Config
	cfg := config{
		port: ":8080",
		db: dbConfig{
			dns: env.GetString("Hadidi-Win-DB-DSN", "host=localhost user=root password=root dbname=hadidiwin sslmode=disable"),
		},
	}

	// Database
	conn, err := pgx.Connect(ctx, cfg.db.dns)
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		panic(err)
	}

	redisClint := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer conn.Close(ctx)

	productRepo := repository.NewProductRepository(conn, redisClint)
	productUsecase := usecase.NewProductUsecase(productRepo)

	r := gin.Default()
	r.Use(middleware.RateLimiter())
	https.NewProductHandler(r, productUsecase)

	log.Printf("Server is Starting at %v", cfg.port)
	r.Run(":%p", cfg.port)
}

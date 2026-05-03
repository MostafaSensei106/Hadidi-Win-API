package cmd

import (
	"time"

	"github.com/jackc/pgx/v5"
)

type application struct {
	config config
	db     *pgx.Conn
}

type config struct {
	port         string
	db           dbConfig
	writeTimeout time.Duration
	readTimeout  time.Duration
	idleTimeout  time.Duration
}

type dbConfig struct {
	dns string
}

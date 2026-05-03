package cmd

import (
	"log"
)

func Execute() {
	cfg := config{
		port: ":8080",
	}
	app := &application{
		config: cfg,
	}
	if err := app.mount(); err != nil {
		log.Fatal(err)
	}
}

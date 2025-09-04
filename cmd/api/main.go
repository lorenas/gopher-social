package main

import (
	"log"

	"github.com/lorenas/gopher-social/internal/env"
	"github.com/lorenas/gopher-social/internal/store"
)

func main() {
	config := config{
		address: env.GetString("ADDR", ":8080"),
	}

	repo := store.NewStorage(nil)

	app := &application{
		config: config,
		repo:   repo,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}

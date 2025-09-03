package main

import (
	"log"

	"github.com/lorenas/gopher-social/internal/env"
)

func main() {
	config := config{
		address: env.GetString("ADDR", ":8080"),
	}
	app := &application{
		config: config,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}

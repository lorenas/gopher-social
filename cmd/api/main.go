package main

import (
	"log"
	"runtime"

	"github.com/lorenas/gopher-social/internal/database"
	"github.com/lorenas/gopher-social/internal/env"
	"github.com/lorenas/gopher-social/internal/store"
)

var (
	databaseDsnFallback = "postgres://admin:adminpassword@localhost/social?sslmode=disable"
)

func main() {
	log.Printf("Starting server with %d CPUs\n", runtime.NumCPU())
	config := config{
		address: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			dsn:          env.GetString("DB_DSN", databaseDsnFallback),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := database.NewDB(config.db.dsn, config.db.maxOpenConns, config.db.maxIdleConns, config.db.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	log.Println("database connection pool established")
	repo := *store.NewStorage(db)

	app := &application{
		config: config,
		repo:   repo,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}

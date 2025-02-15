package main

import (
	"avito-shop/cmd/app"
	"avito-shop/cmd/config"
	"avito-shop/cmd/initDB"
	"avito-shop/seed"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	cfg := config.New()

	db, err := initDB.InitDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	if err := initDB.RunMigrations(db, cfg.Database.Name, "file://migrations"); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	if err := seed.ApplySeeds(db); err != nil {
		log.Fatalf("Failed to apply seeds: %v", err)
	}

	application := app.New(db, cfg.SecretKey)
	if err := application.Run(cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

package main

import (
	"avito-shop/internal/auth/bcrypt"
	"avito-shop/internal/auth/jwt"
	"avito-shop/internal/controller"
	"avito-shop/internal/repository/pg"
	"avito-shop/internal/usecase/implementations"
	"avito-shop/seed"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")
	secretKey := os.Getenv("SECRET_KEY")
	serverPort := ":" + os.Getenv("SERVER_PORT")

	fmt.Printf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable\n", host, port, user, password, dbname)
	fmt.Println("new")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("connection failed: %s", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("connection failed: %s", err.Error())
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to create driver: %s", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		dbname,
		driver,
	)
	if err != nil {
		log.Fatalf("Failed to initialize migrate: %s", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		if !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal(err)
		}
	}

	err = seed.ApplySeeds(db)
	if err != nil {
		log.Fatalf("merch did not upload: %s", err.Error())
	}

	userRepo := pg.NewPgUserRepository(db)
	transactionRepo := pg.NewPgTransactionRepository(db)
	purchaseRepo := pg.NewPgPurchaseRepository(db)
	merchRepo := pg.NewPgMerchRepository(db)
	authService := jwt.NewJWTService(secretKey)
	hashService := bcrypt.NewHashService()

	authUsecase := implementations.NewAuthUsecase(userRepo, authService, hashService)
	userUsecase := implementations.NewUserUsecase(userRepo, transactionRepo, purchaseRepo)
	transactionUsecase := implementations.NewTransactionUseCase(userRepo, transactionRepo)
	purchaseUsecase := implementations.NewPurchaseUsecase(userRepo, merchRepo, purchaseRepo)

	router := controller.SetupRouter(authUsecase, purchaseUsecase, transactionUsecase, userUsecase, authService)

	if err := router.Run(serverPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

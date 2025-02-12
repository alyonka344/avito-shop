package main

import (
	"avito-shop/internal/auth/bcrypt"
	"avito-shop/internal/auth/jwt"
	"avito-shop/internal/model"
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

	userRepository := pg.NewPgUserRepository(db)
	authService := jwt.NewJWTService(secretKey)
	hashService := bcrypt.NewHashService()

	// Инициализируем usecase
	authUsecase := implementations.NewAuthUsecase(userRepository, authService, hashService)

	// Создаем тестового пользователя
	testUser := model.User{
		Username: "alyonka",
		Password: "password",
	}

	// Вызываем создание пользователя
	err = authUsecase.Register(&testUser)
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}

	fmt.Println("User successfully created!")
}

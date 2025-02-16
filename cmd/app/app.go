package app

import (
	"avito-shop/internal/auth/bcrypt"
	"avito-shop/internal/auth/jwt"
	"avito-shop/internal/controller"
	"avito-shop/internal/repository/pg"
	"avito-shop/internal/usecase/usecase_impl"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type App struct {
	Router *gin.Engine
	db     *sqlx.DB
}

func New(db *sqlx.DB, secretKey string) *App {
	repositories := initRepositories(db)
	services := initServices(secretKey)
	useCases := initUseCases(repositories, services)

	router := controller.SetupRouter(
		useCases.auth,
		useCases.purchase,
		useCases.transaction,
		useCases.user,
		services.auth,
	)

	return &App{
		Router: router,
		db:     db,
	}
}

func (a *App) Run(addr string) error {
	return a.Router.Run(addr)
}

type repositories struct {
	user        *pg.PgUserRepository
	transaction *pg.PgTransactionRepository
	purchase    *pg.PgPurchaseRepository
	merch       *pg.PgMerchRepository
}

type services struct {
	auth *jwt.JwtService
	hash *bcrypt.HashService
}

type useCases struct {
	auth        *usecase_impl.AuthUsecase
	user        *usecase_impl.UserUsecase
	transaction *usecase_impl.TransactionUsecase
	purchase    *usecase_impl.PurchaseUsecase
}

func initRepositories(db *sqlx.DB) *repositories {
	return &repositories{
		user:        pg.NewPgUserRepository(db),
		transaction: pg.NewPgTransactionRepository(db),
		purchase:    pg.NewPgPurchaseRepository(db),
		merch:       pg.NewPgMerchRepository(db),
	}
}

func initServices(secretKey string) *services {
	return &services{
		auth: jwt.NewJWTService(secretKey),
		hash: bcrypt.NewHashService(),
	}
}

func initUseCases(r *repositories, s *services) *useCases {
	return &useCases{
		auth:        usecase_impl.NewAuthUsecase(r.user, s.auth, s.hash),
		user:        usecase_impl.NewUserUsecase(r.user, r.transaction, r.purchase),
		transaction: usecase_impl.NewTransactionUseCase(r.user, r.transaction),
		purchase:    usecase_impl.NewPurchaseUsecase(r.user, r.merch, r.purchase),
	}
}

package controller

import (
	"avito-shop/internal/auth"
	"avito-shop/internal/usecase"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	authUsecase usecase.AuthUsecase,
	purchaseUsecase usecase.PurchaseUsecase,
	transactionUsecase usecase.TransactionUsecase,
	userUsecase usecase.UserUsecase,
	authService auth.AuthService) *gin.Engine {
	r := gin.Default()

	r.Use(auth.AuthMiddleware(authService, authUsecase))

	authController := NewAuthController(authUsecase)
	purchaseController := NewPurchaseController(purchaseUsecase)
	transactionController := NewTransactionController(transactionUsecase)
	userController := NewUserController(userUsecase)

	api := r.Group("/api")

	api.POST("/auth", authController.Login)
	api.GET("/info", userController.Info)
	api.POST("/sendCoin", transactionController.SendCoin)
	api.GET("/buy/:item", purchaseController.Buy)

	return r
}

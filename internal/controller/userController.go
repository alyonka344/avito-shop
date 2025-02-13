package controller

import (
	"avito-shop/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *UserController {
	return &UserController{userUsecase: userUsecase}
}

type UserRequest struct {
	MerchName string `json:"merch_name" binding:"required"`
}

type UserResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (pc *UserController) Info(c *gin.Context) {

}

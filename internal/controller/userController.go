package controller

import (
	"avito-shop/internal/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *UserController {
	return &UserController{userUsecase: userUsecase}
}

func (uc *UserController) Info(c *gin.Context) {
	userName, exists := c.Get("userName")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	strUserName, ok := userName.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse username"})
		return
	}

	userInfo, err := uc.userUsecase.GetInfo(strUserName)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user info"})
		return
	}

	c.JSON(http.StatusOK, userInfo)
}

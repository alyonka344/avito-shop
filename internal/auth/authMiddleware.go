package auth

import (
	"avito-shop/internal/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func AuthMiddleware(authService AuthService, authUsecase usecase.AuthUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			var req AuthRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				fmt.Println("token3")
				c.JSON(http.StatusBadRequest, gin.H{"error": "Missing username or password"})
				c.Abort()
				return
			}
			fmt.Println("token")
			user, err := authUsecase.ValidateOrCreateUser(req.Username, req.Password)
			if err != nil {
				fmt.Println("token4")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate or create user"})
				c.Abort()
				return
			}

			token, err := authService.GenerateToken(user)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
				c.Abort()
				return
			}

			c.JSON(http.StatusOK, gin.H{"token": token})
			c.Abort()
			return
		}

		userName, err := authService.ValidateToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userName", userName)
		c.Next()
	}
}

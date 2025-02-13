package auth

import (
	"avito-shop/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware(authService AuthService, authUsecase usecase.AuthUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			username := c.Query("username")
			password := c.Query("password")

			if username == "" || password == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Missing username or password"})
				c.Abort()
				return
			}
			user, err := authUsecase.ValidateOrCreateUser(username, password)
			if err != nil {
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
			return
		}

		userID, err := authService.ValidateToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}

package controller

import (
	"avito-shop/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransactionController struct {
	transactionUsecase usecase.TransactionUsecase
}

func NewTransactionController(transactionUsecase usecase.TransactionUsecase) *TransactionController {
	return &TransactionController{transactionUsecase: transactionUsecase}
}

type TransactionRequest struct {
	Recipient string `json:"recipient" binding:"required"`
	Amount    int    `json:"amount" binding:"required,min=1"`
}

type TransactionResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (tc *TransactionController) SendCoin(c *gin.Context) {
	var req TransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	senderName, exists := c.Get("userName")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	strSenderName, ok := senderName.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse username"})
		return
	}

	err := tc.transactionUsecase.TransferMoney(strSenderName, req.Recipient, req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "transaction failed"})
		return
	}

	c.JSON(http.StatusOK, TransactionResponse{
		Status:  "success",
		Message: "Transaction completed successfully",
	})
}

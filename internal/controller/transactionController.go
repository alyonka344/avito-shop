package controller

import (
	"avito-shop/internal/usecase"
	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	transactionUsecase usecase.TransactionUsecase
}

func NewTransactionController(transactionUsecase usecase.TransactionUsecase) *TransactionController {
	return &TransactionController{transactionUsecase: transactionUsecase}
}

type TransactionRequest struct {
}

type TransactionResponse struct {
}

func (pc *TransactionController) SendCoin(c *gin.Context) {

}

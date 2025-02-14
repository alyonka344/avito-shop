package controller

import (
	"avito-shop/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PurchaseController struct {
	purchaseUsecase usecase.PurchaseUsecase
}

func NewPurchaseController(purchaseUsecase usecase.PurchaseUsecase) *PurchaseController {
	return &PurchaseController{purchaseUsecase: purchaseUsecase}
}

type PurchaseRequest struct {
	MerchName string `json:"merch_name" binding:"required"`
}

type PurchaseResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (pc *PurchaseController) Buy(c *gin.Context) {
	item := c.Param("item")

	userName, exists := c.Get("userName")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	strUserName, ok := userName.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user ID"})
		return
	}

	err := pc.purchaseUsecase.BuyMerch(strUserName, item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to buy merch"})
		return
	}

	c.JSON(http.StatusOK, PurchaseResponse{
		Status:  "success",
		Message: "Purchase completed successfully",
	})
}

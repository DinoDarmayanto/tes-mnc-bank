package handler

import (
	"net/http"
	"tes-mnc-bank/model"
	"tes-mnc-bank/usecase"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	trxUsecase usecase.TransaksiUsecase
}

func (th *TransactionHandler) AddTransaksi(c *gin.Context) {
	// Parsing request body
	var newTransaction model.Transaction
	if err := c.ShouldBindJSON(&newTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	// Add transaction
	err := th.trxUsecase.AddTransaksi(&newTransaction)
	if err != nil {
		// Handle error appropriately, e.g., return a JSON response with error message
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add transaction"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "Transaction added successfully"})
}

func NewTransaksiHandler(srv *gin.Engine, trxUsecase usecase.TransaksiUsecase) *TransactionHandler {
	trxHandler := &TransactionHandler{
		trxUsecase: trxUsecase,
	}

	srv.POST("/Transaksi", trxHandler.AddTransaksi)

	return trxHandler
}

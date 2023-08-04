package handler

import (
	"net/http"
	"strconv"
	"tes-mnc-bank/model"
	"tes-mnc-bank/usecase"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	trxUsecase usecase.TransaksiUsecase
}

func (th *TransactionHandler) GetAllTransaksi(c *gin.Context) {
	// Implementasi kode untuk mendapatkan semua data transaksi dari use case
	transactions, err := th.trxUsecase.GetAllTransaksi()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get transactions"})
		return
	}

	c.JSON(http.StatusOK, transactions)
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
func (th *TransactionHandler) GetTransaksiById(c *gin.Context) {
	// Mendapatkan ID transaksi dari parameter URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	// Implementasi kode untuk mendapatkan data transaksi berdasarkan ID dari use case
	transaction, err := th.trxUsecase.GetTransaksiById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get transaction"})
		return
	}

	// Jika transaksi dengan ID tertentu tidak ditemukan, kembalikan response 404 Not Found
	if transaction == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func NewTransaksiHandler(srv *gin.Engine, trxUsecase usecase.TransaksiUsecase) *TransactionHandler {
	trxHandler := &TransactionHandler{
		trxUsecase: trxUsecase,
	}

	srv.POST("/Transaksi", trxHandler.AddTransaksi)
	srv.GET("/Transaksi", trxHandler.GetAllTransaksi)
	srv.GET("/Transaksi/:id", trxHandler.GetTransaksiById)

	return trxHandler
}

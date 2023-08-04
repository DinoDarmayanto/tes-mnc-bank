// handler/merchants_handler.go

package handler

import (
	"net/http"
	"strconv"
	"strings"
	"tes-mnc-bank/model"
	"tes-mnc-bank/usecase"

	"github.com/gin-gonic/gin"
)

type MerchantHandler struct {
	mctUsecase usecase.MerchantsUsecase
}

func (mctHandler *MerchantHandler) UpdateMerchants(c *gin.Context) {
	// Parsing request body
	var updatedMerchant model.Merchant
	if err := c.ShouldBindJSON(&updatedMerchant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	// Update merchant
	err := mctHandler.mctUsecase.UpdateMerchants(&updatedMerchant)
	if err != nil {
		// Handle specific error cases and return appropriate error messages
		switch {
		case strings.Contains(err.Error(), "invalid merchant data"):
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case strings.Contains(err.Error(), "is not found"):
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update merchant"})
		}
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "Merchant updated successfully"})
}

func (mctHandler *MerchantHandler) AddMerchants(c *gin.Context) {
	// Parsing request body
	var newMerchant model.Merchant
	if err := c.ShouldBindJSON(&newMerchant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	// Add merchant
	err := mctHandler.mctUsecase.AddMerchant(&newMerchant)
	if err != nil {
		// Handle error appropriately, e.g., return a JSON response with error message
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add merchant"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "Merchant added successfully"})
}

func (mctHandler *MerchantHandler) GetAllMerchants(c *gin.Context) {
	// Implementasi kode untuk mendapatkan semua data merchant dari use case
	merchants, err := mctHandler.mctUsecase.GetAllMerchants()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get merchants"})
		return
	}

	c.JSON(http.StatusOK, merchants)
}
func (mctHandler *MerchantHandler) GetMerchantsById(c *gin.Context) {
	// Mendapatkan ID merchant dari parameter URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid merchant ID"})
		return
	}

	// Implementasi kode untuk mendapatkan data merchant berdasarkan ID dari use case
	merchant, err := mctHandler.mctUsecase.GetMerchantsById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get merchant"})
		return
	}

	// Jika merchant dengan ID tertentu tidak ditemukan, kembalikan response 404 Not Found
	if merchant == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Merchant not found"})
		return
	}

	c.JSON(http.StatusOK, merchant)
}
func (mctHandler *MerchantHandler) DeleteMerchants(c *gin.Context) {
	// Mendapatkan ID merchant dari parameter URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid merchant ID"})
		return
	}

	// Implementasi kode untuk menghapus data merchant berdasarkan ID dari use case
	err = mctHandler.mctUsecase.DeleteMerchants(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete merchant"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Merchant deleted successfully"})
}

func NewMerchantsHandler(srv *gin.Engine, mctUsecase usecase.MerchantsUsecase) *MerchantHandler {
	merchantHandler := &MerchantHandler{
		mctUsecase: mctUsecase,
	}

	srv.GET("/merchant", merchantHandler.GetAllMerchants)
	srv.GET("/merchant/:id", merchantHandler.GetMerchantsById)
	srv.DELETE("/merchant/:id", merchantHandler.DeleteMerchants)
	srv.POST("/merchant", merchantHandler.AddMerchants)
	srv.PUT("/merchant", merchantHandler.UpdateMerchants)

	return merchantHandler
}

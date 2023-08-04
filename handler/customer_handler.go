package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"tes-mnc-bank/model"
	"tes-mnc-bank/usecase"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	cstUsecase usecase.CustomerUsecase
}

func (h *CustomerHandler) RegisterCustomer(c *gin.Context) {
	// Parsing request body
	// Parsing request body
	var user model.Customer
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err) // Tambahkan baris ini untuk mencetak kesalahan
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	// Validasi dan registrasi pengguna
	if err := h.cstUsecase.RegisterCustomer(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan response dengan data pengguna yang terdaftar
	c.JSON(http.StatusOK, user)
}
func (h *CustomerHandler) DeleteCustomer(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMassage": "Id tidak boleh kosong",
		})
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMassage": "Id harus angka",
		})
		return
	}

	err = h.cstUsecase.DeleteCustomer(id)
	if err != nil {
		fmt.Printf("customerHandlerImpl.DeleteCustomer(() : %v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMassage": "Terjadi kesalahan ketika menghapus data customer",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"Massage": "Customer telah di hapus",
		"ID":      id,
	})
}

func (h *CustomerHandler) LoginCustomer(ctx *gin.Context) {
	// Parsing request body
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	// Login user
	token, err := h.cstUsecase.Login(credentials.Email, credentials.Password)
	if err != nil {
		// Mengembalikan response dengan pesan error yang lebih umum
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Mengembalikan response dengan token
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func NewCustomerHandler(srv *gin.Engine, cstUsecase usecase.CustomerUsecase) *CustomerHandler {
	customerHandler := &CustomerHandler{
		cstUsecase: cstUsecase,
	}

	srv.POST("/register", customerHandler.RegisterCustomer)
	srv.POST("/login", customerHandler.LoginCustomer)

	// Gunakan metode GET untuk menghapus pelanggan
	srv.GET("/customer/:id", customerHandler.DeleteCustomer)

	return customerHandler
}

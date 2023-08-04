package handler

import (
	"net/http"
	"tes-mnc-bank/usecase"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	LoginUsecase usecase.LoginUsecase
}

func (l *LoginHandler) Logout(c *gin.Context) {
	err := l.LoginUsecase.Logout(c) // Sertakan konteks c saat memanggil Logout
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Logout successful",
	})
}

func NewLoginHandler(router *gin.Engine, loginUsecase usecase.LoginUsecase) {
	// Inisialisasi cookie store
	store := cookie.NewStore([]byte("secret-key")) // Ganti dengan kunci rahasia yang lebih kuat

	// Set pengaturan cookie store
	store.Options(sessions.Options{
		HttpOnly: true,
		Secure:   true,
	})

	//Inisialisasi handler dengan penggunaan sesi
	loginHandler := &LoginHandler{
		LoginUsecase: loginUsecase,
	}

	router.Use(sessions.Sessions("session-name", store))

	// router.POST("/login", loginHandler.Login)
	router.POST("/logout", loginHandler.Logout)
}

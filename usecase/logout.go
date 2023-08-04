package usecase

import (
	"tes-mnc-bank/repository"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

type LoginUsecase interface {
	// Login(email, password string) (string, error)
	Logout(ctx *gin.Context) error
}

type loginUsecase struct {
	loginRepo repository.CustomerRepo
}

func (lu *loginUsecase) Logout(ctx *gin.Context) error {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
	return nil
}

func NewLoginUsecase(loginRepo repository.CustomerRepo) LoginUsecase {
	return &loginUsecase{
		loginRepo: loginRepo,
	}
}

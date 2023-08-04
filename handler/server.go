package handler

import (
	"tes-mnc-bank/config"
	"tes-mnc-bank/manager"
	"tes-mnc-bank/middleware"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Run()
}

type server struct {
	usecaseManager manager.UsecaseManager
	engine         *gin.Engine
}

func (s *server) Run() {
	s.engine.Use(middleware.LoggerMiddleware())

	NewCustomerHandler(s.engine, s.usecaseManager.GetCustomerUsecase())
	NewMerchantsHandler(s.engine, s.usecaseManager.GetMerchantsUsecase())
	NewTransaksiHandler(s.engine, s.usecaseManager.GetTransaksiUsecase())
	NewLoginHandler(s.engine, s.usecaseManager.GetLoginUsecase())

	s.engine.Run(":8080")
}

func NewServer() Server {
	c, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	infra := manager.NewInfraManager(c)
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewUsecaseManager(repo)

	engine := gin.Default()

	return &server{
		usecaseManager: usecase,
		engine:         engine,
	}
}

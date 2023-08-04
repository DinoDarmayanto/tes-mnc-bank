package manager

import (
	"sync"
	"tes-mnc-bank/usecase"
)

type UsecaseManager interface {
	GetCustomerUsecase() usecase.CustomerUsecase
	GetMerchantsUsecase() usecase.MerchantsUsecase
	GetTransaksiUsecase() usecase.TransaksiUsecase
}

type usecaseManager struct {
	repoManager RepoManager
	cstUsecase  usecase.CustomerUsecase
	mctUsecase  usecase.MerchantsUsecase
	trxUsecase  usecase.TransaksiUsecase
}

var onceLoadCustomerUsecase sync.Once
var onceLoadMerchantsUsecase sync.Once
var onceLoadTransaksiUsecase sync.Once

func (um *usecaseManager) GetCustomerUsecase() usecase.CustomerUsecase {
	onceLoadCustomerUsecase.Do(func() {
		um.cstUsecase = usecase.NewCustomerUseCase(um.repoManager.GetCustomerRepo())
	})
	return um.cstUsecase
}
func (um *usecaseManager) GetMerchantsUsecase() usecase.MerchantsUsecase {
	onceLoadMerchantsUsecase.Do(func() {
		um.mctUsecase = usecase.NewMerchantsUseCase(um.repoManager.GetMerchantsRepo())
	})
	return um.mctUsecase
}

func (um *usecaseManager) GetTransaksiUsecase() usecase.TransaksiUsecase {
	onceLoadTransaksiUsecase.Do(func() {
		um.trxUsecase = usecase.NewTransaksiUseCase(um.repoManager.GetTransaksiRepo())
	})
	return um.trxUsecase
}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}

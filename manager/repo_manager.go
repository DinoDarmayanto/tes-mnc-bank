package manager

import (
	"sync"
	"tes-mnc-bank/repository"
)

type RepoManager interface {
	GetCustomerRepo() repository.CustomerRepo
	GetMerchantsRepo() repository.MerchantsRepo
	GetTransaksiRepo() repository.TransaksiRepo
}

type repoManager struct {
	infraManager InfraManager
	cstRepo      repository.CustomerRepo
	mctRepo      repository.MerchantsRepo
	trxRepo      repository.TransaksiRepo
}

var onceLoadCustomerRepo sync.Once
var onceLoadMerchantsRepo sync.Once
var onceLoadTransaksiRepo sync.Once

func (rm *repoManager) GetCustomerRepo() repository.CustomerRepo {
	onceLoadCustomerRepo.Do(func() {
		rm.cstRepo = repository.NewCustomerRepo(rm.infraManager.GetDB())
	})
	return rm.cstRepo
}

func (rm *repoManager) GetMerchantsRepo() repository.MerchantsRepo {
	onceLoadMerchantsRepo.Do(func() {
		rm.mctRepo = repository.NewMerchantsRepo(rm.infraManager.GetDB())
	})
	return rm.mctRepo
}
func (rm *repoManager) GetTransaksiRepo() repository.TransaksiRepo {
	onceLoadTransaksiRepo.Do(func() {
		rm.trxRepo = repository.NewTransaksiRepo(rm.infraManager.GetDB())
	})
	return rm.trxRepo
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}

package usecase

import (
	"fmt"
	"tes-mnc-bank/model"
	"tes-mnc-bank/repository"
)

type TransaksiUsecase interface {
	AddTransaksi(merchant *model.Transaction) error
	// GetTransaksiById(id int) (*model.Merchant, error)
	// GetAllMerchants() ([]model.Merchant, error)
	// DeleteMerchants(id int) error
}

type TransactionUsecaseImpl struct {
	trxRepo repository.TransaksiRepo
}

func (uc *TransactionUsecaseImpl) AddTransaksi(trx *model.Transaction) error {
	// Add validation logic here if needed

	// Call the AddTransaksi method from the repository
	err := uc.trxRepo.AddTransaksi(trx)
	if err != nil {
		// You can add additional context to the error if needed
		return fmt.Errorf("TransactionUsecaseImpl.AddTransaksi: failed to add transaction: %w", err)
	}

	return nil
}

func NewTransaksiUseCase(trxRepo repository.TransaksiRepo) TransaksiUsecase {
	return &TransactionUsecaseImpl{
		trxRepo: trxRepo,
	}
}

package usecase

import (
	"fmt"
	"tes-mnc-bank/model"
	"tes-mnc-bank/repository"
)

type TransaksiUsecase interface {
	AddTransaksi(merchant *model.Transaction) error
	GetTransaksiById(id int) (*model.Transaction, error)
	GetAllTransaksi() ([]model.Transaction, error)
}

type TransactionUsecaseImpl struct {
	trxRepo repository.TransaksiRepo
}

func (uc *TransactionUsecaseImpl) GetAllTransaksi() ([]model.Transaction, error) {
	// Call the GetAllTransaksi method from the repository
	transactions, err := uc.trxRepo.GetAllTransaksi()
	if err != nil {
		// You can add additional context to the error if needed
		return nil, fmt.Errorf("TransactionUsecaseImpl.GetAllTransaksi: failed to get all transactions: %w", err)
	}

	return transactions, nil
}
func (uc *TransactionUsecaseImpl) GetTransaksiById(id int) (*model.Transaction, error) {
	// Call the GetTransaksiById method from the repository
	transaction, err := uc.trxRepo.GetTransaksiById(id)
	if err != nil {
		// You can add additional context to the error if needed
		return nil, fmt.Errorf("TransactionUsecaseImpl.GetTransaksiById: failed to get transaction by ID: %w", err)
	}

	return transaction, nil
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

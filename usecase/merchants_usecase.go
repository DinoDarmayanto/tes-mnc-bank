package usecase

import (
	"fmt"
	"log"
	"tes-mnc-bank/apperror"
	"tes-mnc-bank/model"
	"tes-mnc-bank/repository"
)

type MerchantsUsecase interface {
	GetAllMerchants() ([]model.Merchant, error)
	AddMerchant(mct *model.Merchant) error
	GetMerchantsById(id int) (*model.Merchant, error)
	DeleteMerchants(id int) error
	UpdateMerchants(merchant *model.Merchant) error
}

type merchantsUsecaseImpl struct {
	mctRepo repository.MerchantsRepo
}

func (mctUsecase *merchantsUsecaseImpl) UpdateMerchants(mct *model.Merchant) error {
	// Pastikan data merchant valid sebelum melakukan operasi di database
	if mct.Name == "" && mct.AccountNumber == "" {
		return fmt.Errorf("merchantsUsecaseImpl.UpdateMerchants(): both name and account number are empty")
	} else if mct.Name == "" {
		return fmt.Errorf("merchantsUsecaseImpl.UpdateMerchants(): invalid merchant data, name is empty")
	} else if mct.AccountNumber == "" {
		return fmt.Errorf("merchantsUsecaseImpl.UpdateMerchants(): invalid merchant data, account number is empty")
	}

	mctDb, err := mctUsecase.mctRepo.GetMerchantByAccountNumber(mct.AccountNumber)
	if err != nil {
		return fmt.Errorf("merchantsUsecaseImpl.UpdateMerchants(): error fetching merchant by account number: %w", err)
	}

	if mctDb == nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("Account number %v is not found!", mct.AccountNumber),
		}
	}

	mctDb.Name = mct.Name
	mctDb.AccountNumber = mct.AccountNumber

	err = mctUsecase.mctRepo.UpdateMerchants(mctDb)
	if err != nil {
		// Log error detail sebelum mengembalikan error
		log.Printf("merchantsUsecaseImpl.UpdateMerchants(): error updating merchant: %v", err)
		return fmt.Errorf("merchantsUsecaseImpl.UpdateMerchants(): failed to update merchant: %w", err)
	}

	return nil
}

func (mctUsecase *merchantsUsecaseImpl) GetAllMerchants() ([]model.Merchant, error) {
	return mctUsecase.mctRepo.GetAllMerchants()
}
func (mctUsecase *merchantsUsecaseImpl) DeleteMerchants(id int) error {
	err := mctUsecase.mctRepo.DeleteMerchants(id)
	if err != nil {
		return fmt.Errorf("error on merchantsUsecaseImpl.DeleteMerchants(): %w", err)
	}
	return nil
}

func (mctUsecase *merchantsUsecaseImpl) GetMerchantsById(id int) (*model.Merchant, error) {
	return mctUsecase.mctRepo.GetMerchantsById(id)
}

func (mctUsecase *merchantsUsecaseImpl) AddMerchant(mct *model.Merchant) error {
	// Pastikan data merchant valid sebelum melakukan operasi di database
	if mct.Name == "" || mct.AccountNumber == "" {
		return fmt.Errorf("merchantsUsecaseImpl.AddMerchant(): invalid merchant data")
	}

	mctDb, err := mctUsecase.mctRepo.GetMerchantByAccountNumber(mct.AccountNumber)
	if err != nil {
		return fmt.Errorf("merchantsUsecaseImpl.AddMerchant(): error fetching merchant by account number: %w", err)
	}

	if mctDb != nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("Account number %v is already registered!", mct.AccountNumber),
		}
	}

	err = mctUsecase.mctRepo.AddMerchant(mct)
	if err != nil {
		// Log error detail sebelum mengembalikan error
		log.Printf("merchantsUsecaseImpl.AddMerchant(): error adding merchant: %v", err)
		return fmt.Errorf("merchantsUsecaseImpl.AddMerchant(): failed to add merchant: %w", err)
	}

	return nil
}

func NewMerchantsUseCase(mctRepo repository.MerchantsRepo) MerchantsUsecase {
	return &merchantsUsecaseImpl{
		mctRepo: mctRepo,
	}
}

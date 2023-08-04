package repository

import (
	"database/sql"
	"fmt"
	"tes-mnc-bank/model"
)

type MerchantsRepo interface {
	AddMerchant(merchant *model.Merchant) error
	GetMerchantsById(id int) (*model.Merchant, error)
	GetMerchantByAccountNumber(accountNumber string) (*model.Merchant, error)
	GetAllMerchants() ([]model.Merchant, error)
	DeleteMerchants(id int) error
	UpdateMerchants(merchant *model.Merchant) error
}

type MerchantsRepoImpl struct {
	db *sql.DB
}

func (mctRepo *MerchantsRepoImpl) UpdateMerchants(merchant *model.Merchant) error {
	qry := "UPDATE merchants SET name = $1, account_number = $2 WHERE id = $3"

	// Eksekusi query menggunakan prepared statement untuk menghindari SQL injection
	_, err := mctRepo.db.Exec(qry, merchant.Name, merchant.AccountNumber, merchant.Id)
	if err != nil {
		return fmt.Errorf("error on MerchantsRepoImpl.UpdateMerchants(): %w", err)
	}

	return nil
}

func (mctRepo *MerchantsRepoImpl) DeleteMerchants(id int) error {
	qry := "DELETE FROM merchants WHERE id = $1"

	_, err := mctRepo.db.Exec(qry, id)
	if err != nil {
		return fmt.Errorf("error on MerchantsRepoImpl.DeleteMerchants(): %w", err)
	}

	return nil
}

func (mctRepo *MerchantsRepoImpl) AddMerchant(merchant *model.Merchant) error {
	qry := "INSERT INTO merchants (name, account_number) VALUES ($1, $2) RETURNING id"

	// Eksekusi query menggunakan prepared statement untuk menghindari SQL injection
	err := mctRepo.db.QueryRow(qry, merchant.Name, merchant.AccountNumber).Scan(&merchant.Id)
	if err != nil {
		return fmt.Errorf("error on MerchantsRepoImpl.AddMerchant(): %w", err)
	}

	return nil
}

func (mctRepo *MerchantsRepoImpl) GetMerchantsById(id int) (*model.Merchant, error) {
	qry := "SELECT id, name, account_number, registered_at FROM merchants WHERE id = $1"

	row := mctRepo.db.QueryRow(qry, id)

	mct := &model.Merchant{}
	err := row.Scan(&mct.Id, &mct.Name, &mct.AccountNumber, &mct.RegisteredAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on MerchantsRepoImpl.GetMerchantsById(): %w", err)
	}

	return mct, nil
}

// MerchantsRepoImpl
func (mctRepo *MerchantsRepoImpl) GetMerchantByAccountNumber(accountNumber string) (*model.Merchant, error) {
	qry := "SELECT id, name, account_number, registered_at FROM merchants WHERE account_number = $1"

	row := mctRepo.db.QueryRow(qry, accountNumber)

	mct := &model.Merchant{}
	err := row.Scan(&mct.Id, &mct.Name, &mct.AccountNumber, &mct.RegisteredAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on MerchantsRepoImpl.GetMerchantByAccountNumber(): %w", err)
	}

	return mct, nil
}

func (mctRepo *MerchantsRepoImpl) GetAllMerchants() ([]model.Merchant, error) {
	qry := "SELECT id, name, account_number, registered_at FROM merchants"

	rows, err := mctRepo.db.Query(qry)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on customerRepoImpl.GetAllCustomer(): %w", err)
	}
	defer rows.Close()

	arrMct := []model.Merchant{}
	for rows.Next() {
		mct := model.Merchant{}
		err := rows.Scan(&mct.Id, &mct.Name, &mct.AccountNumber, &mct.RegisteredAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		arrMct = append(arrMct, mct)
	}
	return arrMct, nil
}

func NewMerchantsRepo(db *sql.DB) MerchantsRepo {
	return &MerchantsRepoImpl{
		db: db,
	}
}

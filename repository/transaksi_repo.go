package repository

import (
	"database/sql"
	"fmt"
	"tes-mnc-bank/model"
)

type TransaksiRepo interface {
	AddTransaksi(merchant *model.Transaction) error
	// GetTransaksiById(id int) (*model.Merchant, error)
	// GetAllMerchants() ([]model.Merchant, error)
	// DeleteMerchants(id int) error
}

type TransaksiRepoImpl struct {
	db *sql.DB
}

func (trxRepo *TransaksiRepoImpl) AddTransaksi(Trx *model.Transaction) error {
	qry := "INSERT INTO transactions (customer_id, merchant_id, amount, transaction_time) VALUES ($1, $2, $3, $4) RETURNING id"

	// Eksekusi query menggunakan prepared statement untuk menghindari SQL injection
	err := trxRepo.db.QueryRow(qry, Trx.CustomerID, Trx.MerchantID, Trx.Amount, Trx.TransactionTime).Scan(&Trx.ID)
	if err != nil {
		return fmt.Errorf("error on MerchantsRepoImpl.AddTransaksi(): %w", err)
	}

	return nil
}

func NewTransaksiRepo(db *sql.DB) TransaksiRepo {
	return &TransaksiRepoImpl{
		db: db,
	}
}

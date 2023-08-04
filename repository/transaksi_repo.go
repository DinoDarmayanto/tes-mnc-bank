package repository

import (
	"database/sql"
	"fmt"
	"tes-mnc-bank/model"
)

type TransaksiRepo interface {
	AddTransaksi(merchant *model.Transaction) error
	GetTransaksiById(id int) (*model.Transaction, error)
	GetAllTransaksi() ([]model.Transaction, error)
}

type TransaksiRepoImpl struct {
	db *sql.DB
}

func (trxRepo *TransaksiRepoImpl) GetTransaksiById(id int) (*model.Transaction, error) {
	qry := "SELECT id, customer_id, merchant_id, amount, transaction_time FROM transactions WHERE id = $1"

	row := trxRepo.db.QueryRow(qry, id)

	trx := &model.Transaction{}
	err := row.Scan(&trx.ID, &trx.CustomerID, &trx.MerchantID, &trx.Amount, &trx.TransactionTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on TransaksiRepoImpl.GetTransaksiById(): %w", err)
	}

	return trx, nil
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
func (trxRepo *TransaksiRepoImpl) GetAllTransaksi() ([]model.Transaction, error) {
	qry := "SELECT id, customer_id, merchant_id, amount, transaction_time FROM transactions"

	rows, err := trxRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("error on TransaksiRepoImpl.GetAllTransaksi(): %w", err)
	}
	defer rows.Close()

	transactions := []model.Transaction{}
	for rows.Next() {
		trx := model.Transaction{}
		err := rows.Scan(&trx.ID, &trx.CustomerID, &trx.MerchantID, &trx.Amount, &trx.TransactionTime)
		if err != nil {
			return nil, fmt.Errorf("error on TransaksiRepoImpl.GetAllTransaksi(): %w", err)
		}
		transactions = append(transactions, trx)
	}

	return transactions, nil
}

func NewTransaksiRepo(db *sql.DB) TransaksiRepo {
	return &TransaksiRepoImpl{
		db: db,
	}
}

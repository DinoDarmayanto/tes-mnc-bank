package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"tes-mnc-bank/model"
	"time"
)

type CustomerRepo interface {
	RegistCustomer(*model.Customer) error
	GetCustomerByUsername(string) (*model.Customer, error)
	GetCustomerByEmail(email string) (*model.Customer, error)
	DeleteCustomer(id int) error
}

type customerRepoImpl struct {
	db *sql.DB
}

func (cstRepo *customerRepoImpl) RegistCustomer(cst *model.Customer) error {
	qry := "INSERT INTO customers (username, password, email, registered_at) VALUES ($1, $2, $3, $4) RETURNING id"

	// Eksekusi query menggunakan database yang sesuai (ganti sesuai dengan database yang Anda gunakan)
	err := cstRepo.db.QueryRow(qry, cst.Username, cst.Password, cst.Email, time.Now()).Scan(&cst.ID)
	if err != nil {
		return fmt.Errorf("error on customerRepoImpl.RegistCustomer(): %w", err)
	}

	return nil
}

func (cst *customerRepoImpl) GetCustomerByUsername(username string) (*model.Customer, error) {
	selectStatement := "SELECT id, username, password, email, registered_at FROM customers WHERE username = $1"

	row := cst.db.QueryRow(selectStatement, username)

	customer := &model.Customer{}
	err := row.Scan(&customer.ID, &customer.Username, &customer.Password, &customer.Email, &customer.RegisteredAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return customer, nil
}
func (cst *customerRepoImpl) GetCustomerByEmail(email string) (*model.Customer, error) {
	selectStatment := "SELECT id, username,  password, email ,registered_at FROM customers WHERE email = $1"

	row := cst.db.QueryRow(selectStatment, email)

	customer := &model.Customer{}
	err := row.Scan(&customer.ID, &customer.Username, &customer.Password, &customer.Email, &customer.RegisteredAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return customer, nil
}

func (cstRepo *customerRepoImpl) DeleteCustomer(id int) error {
	qry := "DELETE FROM customers WHERE id = $1"

	// Eksekusi query menggunakan database yang sesuai (ganti sesuai dengan database yang Anda gunakan)
	_, err := cstRepo.db.Exec(qry, id)
	if err != nil {
		return fmt.Errorf("error on customerRepoImpl.DeleteCustomerByID(): %w", err)
	}

	return nil
}

func NewCustomerRepo(db *sql.DB) CustomerRepo {
	return &customerRepoImpl{
		db: db,
	}
}

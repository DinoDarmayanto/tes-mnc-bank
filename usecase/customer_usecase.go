package usecase

import (
	"errors"
	"fmt"
	"log"
	"tes-mnc-bank/model"
	"tes-mnc-bank/repository"
	"tes-mnc-bank/utils"
	"tes-mnc-bank/utils/authutil"
	"time"

	"golang.org/x/crypto/bcrypt" // Import bcrypt package
)

type CustomerUsecase interface {
	RegisterCustomer(*model.Customer) error
	DeleteCustomer(id int) error
	Login(email, password string) (string, error)
}

type customerUsecaseImpl struct {
	cstRepo repository.CustomerRepo
}

func (cst *customerUsecaseImpl) RegisterCustomer(customer *model.Customer) error {
	// Validasi kredensial
	existingCustomer, err := cst.cstRepo.GetCustomerByUsername(customer.Username)
	if err == nil && existingCustomer != nil {
		return errors.New("username is already taken")
	}

	// Validasi email unik
	existingCustomer, err = cst.cstRepo.GetCustomerByEmail(customer.Email)
	if err == nil && existingCustomer != nil {
		return errors.New("email is already taken")
	}

	// Validasi password keamanan
	if !utils.IsValidPassword(customer.Password) {
		return errors.New("password is not strong enough")
	}

	// Validasi data pengguna
	if !utils.IsValidEmail(customer.Email) {
		return errors.New("invalid email address")
	}

	// Generate hash password
	passHash, err := GeneratePasswordHash(customer.Password)
	if err != nil {
		return fmt.Errorf("failed to generate password hash: %w", err)
	}
	customer.Password = passHash

	// Set nilai default
	customer.RegisteredAt = time.Now()

	// Buat pengguna baru
	err = cst.cstRepo.RegistCustomer(customer)
	if err != nil {
		return err
	}

	return nil
}

func GeneratePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
func (cstUsecase *customerUsecaseImpl) DeleteCustomer(id int) error {
	return cstUsecase.cstRepo.DeleteCustomer(id)

}

func (cstUsecase *customerUsecaseImpl) Login(email, password string) (string, error) {
	// Cari pengguna berdasarkan email
	log.Printf("Trying to login with email: %s", email)

	customer, err := cstUsecase.cstRepo.GetCustomerByEmail(email)
	if err != nil {
		log.Printf("Failed to fetch customer: %v", err)
		return "", fmt.Errorf("login failed: %w", err)
	}

	if customer == nil {
		log.Printf("Customer not found with email: %s", email)
		return "", errors.New("login failed: email not found")
	}

	// Verifikasi password
	log.Println("Verifying password...")
	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(password))
	if err != nil {
		log.Printf("Password verification failed: %v", err)
		return "", errors.New("login failed: incorrect password")
	}

	// Generate token
	log.Println("Generating token...")
	token, err := authutil.GenerateToken(customer)
	if err != nil {
		log.Printf("Failed to generate token: %v", err)
		return "", fmt.Errorf("login failed: %w", err)
	}

	log.Println("Login successful!")
	return token, nil
}

func NewCustomerUseCase(cstRepo repository.CustomerRepo) CustomerUsecase {
	return &customerUsecaseImpl{
		cstRepo: cstRepo,
	}
}

package common

import (
	"errors"

	"github.com/joho/godotenv"
)

func LoadFileEnv(file string) error {
	err := godotenv.Load(file)
	if err != nil {
		return errors.New("failed to load .env file")
	}
	return nil
}

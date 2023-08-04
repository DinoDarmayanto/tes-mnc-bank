package authutil

import (
	"fmt"
	"tes-mnc-bank/model"
	"time"

	"github.com/golang-jwt/jwt"
)

const TokenKey = "a12jasdasb^&*adjsabKJBadASJasb"

type JwtClaims struct {
	Username string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(customer *model.Customer) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtClaims{
		Username: customer.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	})

	// Menandatangani token dengan kunci rahasia
	tokenString, err := token.SignedString([]byte(TokenKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyAccessToken(tokenString string) (string, error) {
	claim := &JwtClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(TokenKey), nil
	})
	if err != nil {
		return "", fmt.Errorf("VerifyAccessToken: %w", err)
	}
	if !token.Valid {
		return "", fmt.Errorf("VerifyAccessToken: Invalid token")
	}
	return claim.Username, nil
}

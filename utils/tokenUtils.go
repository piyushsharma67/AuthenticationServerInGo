package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(email string) (string, error) {
	claims:= CustomClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10000 * time.Hour).Unix(),
		},
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte("test"))
	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateToken(reqtoken string) (bool,string,string){
	if len(reqtoken)!=0 {
		claims := &CustomClaims{}
		token, err := jwt.ParseWithClaims(reqtoken, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Invalid Token")
			}
			return []byte("test"), nil
		})
		if token.Valid {
			return true, claims.UserId,claims.UserId
		}
		return false, err.Error(),""
	}
	return false, "token not provided",""
}
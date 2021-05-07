package apiserver

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type Claims struct {
	jwt.StandardClaims
	Login string `json:"login"`
}

func GenerateToken(login string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(14 * time.Hour)),
			IssuedAt:  jwt.At(time.Now()),
		},
		Login: login,
	},
	)

	return token.SignedString([]byte(os.Getenv("SECRET_JWT_KEY")))
}

func parseToken(signedToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signed method %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_JWT_KEY")), nil
	})

	if err != nil {
		return &Claims{}, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok {
		return &Claims{}, errors.New("Couldn't parse claims")
	}

	if claims.ExpiresAt.Local().Unix() < jwt.At(time.Now()).Local().Unix() {
		return &Claims{}, errors.New("JWT is expired")
	}

	return claims, nil
}

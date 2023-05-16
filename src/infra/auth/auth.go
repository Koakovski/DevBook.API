package auth

import (
	"devbook-api/src/infra/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userId uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString(config.SecretKey)
}

func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)

	token, err := jwt.Parse(tokenString, getVerificationKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

func extractToken(r *http.Request) string {
	authorizationString := r.Header.Get("Authorization")

	token := strings.Split(authorizationString, " ")
	if len(token) == 2 {
		return token[1]
	}

	return ""
}

func getVerificationKey(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signature method. %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}

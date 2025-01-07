package auth

import (
	"devbookAPI/src/config"
	"devbookAPI/src/model"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenToken: genenerates authentication token
func GenToken(user model.User) (string, error) {
	claims := jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 48).Unix(),
	}

	token, err := jwt.
		NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString(config.SECRET)

	if err != nil {
		return "", err
	}

	return token, nil
}

// ValidToken: Check if given token is valid
func ValidToken(r *http.Request) error {
	token, err := getToken(r)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !(ok && token.Valid) {
		return errors.New("token format is invalid")
	}

	return nil
}

// GetAuthenticatedId: Check if given token is valid
func GetAuthenticatedId(r *http.Request) (int64, error) {
	token, err := getToken(r)
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id, err := strconv.ParseInt(fmt.Sprintf("%.0f", claims["sub"]), 10, 64)
		if err != nil {
			return 0, err
		}

		return id, nil
	}

	return 0, errors.New("token format is invalid")
}

// CheckUserPermision: check if current user is
func CheckUserPermision(entityId int64, authId int64) error {
	if authId != entityId {
		return errors.New("authenticated user cannot use this entity")
	}

	return nil
}

func getToken(r *http.Request) (token *jwt.Token, err error) {
	tokenString, err := getTokenString(r)
	if err != nil {
		return
	}

	token, err = jwt.Parse(tokenString, getSecretKey)
	if err != nil {
		return
	}

	return
}

func getTokenString(r *http.Request) (string, error) {
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")

	if len(tokenString) != 2 && tokenString[0] != "Bearer" {
		return "", fmt.Errorf("token format is invalid")
	}

	return tokenString[1], nil
}

func getSecretKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signature method: %v", token.Header["alg"])
	}

	return config.SECRET, nil
}

package auth

import (
	"devbookAPI/src/config"
	"devbookAPI/src/model"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Genenerates user authentication token
func GenUserToken(user model.User) (string, error) {
	claims := jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 48).Unix(),
	}

	token, err := jwt.
		NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString(config.SECRET)

	fmt.Printf("err: %v\n", err)

	if err != nil {
		return "", err
	}

	return token, nil
}

// func GetClains(tokenString string) (*jwt.Claims, error) {
// 	doCheck := func(token *jwt.Token) (interface{}, error) {
// 		if token.Method != jwt.SigningMethodHS256 {
// 			return nil, fmt.Errorf("unexpected signature method: %v", token.Header["alg"])
// 		}

// 		return config.SECRET, nil
// 	}

// 	token, err := jwt.Parse(tokenString, doCheck)
// 	if ; err != nil {
// 		return nil, err
// 	}

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		fmt.Println("Token válido!")
// 		fmt.Println("Claims:", claims)
// 	}

// 	return token, nil

// }

package authentication

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/satimoto/go-datastore/db"
)

func SignToken(user db.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = user.ID
	
	jwtWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtWithClaims.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyToken(tokenString string) (bool, jwt.MapClaims) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return false, jwt.MapClaims{}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, claims
	} 

	return false, jwt.MapClaims{}
}
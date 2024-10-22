package helper

import (
	"fmt"
	"log"
	"time"

	"github.com/RaflyDioniswaraPramono/my-portofolio/configs"
	"github.com/golang-jwt/jwt/v4"
)

type JwtPayload struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	RoleId   int    `json:"role_id"`
	jwt.RegisteredClaims
}

var jwtSecretKey = configs.LoadConfig("development", "JWT_SECRET_KEY")

func GenerateToken(payloads JwtPayload) string {
	claims := jwt.MapClaims{
		"email":    payloads.Email,
		"username": payloads.Username,
		"role_id":  payloads.RoleId,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err := token.SignedString(jwtSecretKey)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return signedString
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return err
	}

	return nil
}

func DecodeToken(tokenString string) (*JwtPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtPayload{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JwtPayload); ok && token.Valid {
		if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
			return nil, err
		}

		return claims, nil
	}

	return nil, err
}

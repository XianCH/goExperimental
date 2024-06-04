package config

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var JwtSecret = []byte("x14n")

type Claims struct {
	Username string `json:"username"`
	UUID     string `json:"uuid"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		Username: username,
		UUID:     uuid.NewString(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(nowTime),
			NotBefore: jwt.NewNumericDate(nowTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(JwtSecret)
}

func RefreshToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		nowTime := time.Now()
		expireTime := nowTime.Add(3 * time.Hour)
		claims.ExpiresAt = jwt.NewNumericDate(expireTime)
		claims.IssuedAt = jwt.NewNumericDate(nowTime)
		claims.NotBefore = jwt.NewNumericDate(nowTime)
		newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return newToken.SignedString(JwtSecret)
	}

	return "", err

}

// ValidToken 验证 JWT
func ValidToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid toekn")
}

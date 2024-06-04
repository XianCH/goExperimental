package jwtexp

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MyCliams struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 2

var MySecurt = []byte("x14nch911")

func GenToken(username string) (string, error) {
	c := MyCliams{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "my-project",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(MySecurt)
}

func ParseToken(tokenString string) (*MyCliams, error) {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyCliams{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecurt, nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if claims, ok := token.Claims.(*MyCliams); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")

}

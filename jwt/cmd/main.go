package main

import (
	"fmt"

	"github.com/x14n/goExperimental/jwt/config"
)

func main() {
	token, err := config.GenerateToken("x14n")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(token)
	token, err = config.RefreshToken(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(token)

	claims, err := config.ValidToken(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(claims)
}

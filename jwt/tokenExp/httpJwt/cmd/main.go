package main

import (
	"github.com/x14n/goExperimental/gorm/gorm_database"
	httpjwt "github.com/x14n/goExperimental/jwt/tokenExp/httpJwt"
)

func main() {
	gorm_database.InitDB()
	router := httpjwt.InitRouter()
	router.Run(":8080")
}

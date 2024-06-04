package main

import "github.com/x14n/goExperimental/jwt/server"

func main() {
	server := server.NewRouter()
	server.Run(":12345")

}

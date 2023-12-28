package main

import (
	"github.com/gin-gonic/gin"
	"github.com/x14n/goExperimental/gin/database"
	"github.com/x14n/goExperimental/gin/src"
)

func main() {
	router := gin.Default()

	database.InitDB()
	v1 := router.Group("/v1")
	src.AddUserRouter(v1)
	router.Run(":8000")
}

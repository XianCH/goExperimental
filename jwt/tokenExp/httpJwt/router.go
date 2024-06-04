package httpjwt

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	file, _ := os.Create("logs/app/log")
	gin.DefaultWriter = io.MultiWriter(file)
	router := gin.Default()
	router.Use(gin.Recovery(), gin.Logger())
	//登录注册
	router.POST("/login", Login)
	// router.POST("/register", Register)
	//用户相关
	userRoute := router.Group("user")
	userRoute.Use(JWTAUTH()) //这里使用Use,jwtAuth就行
	userRoute.GET("/list", UserList)

	return router
}

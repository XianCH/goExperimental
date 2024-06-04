package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/x14n/goExperimental/jwt/config"
	"github.com/x14n/goExperimental/jwt/middleware"
)

func NewRouter() *gin.Engine {

	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.POST("/Login", Login)

	auth := router.Group("/auth")
	auth.Use(middleware.JWT())
	{
		auth.GET("/protected", ProtectedEndpoint)
	}
	return router
}

func Login(c *gin.Context) {
	// 1. 从请求中提取用户名
	username := c.PostForm("username")

	// 2. 验证用户名（可选）
	// 如果用户名无效，返回错误响应
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid username",
		})
		return
	}

	// 3. 使用给定的用户名生成 JWT token
	token, err := config.GenerateToken(username)
	if err != nil {
		// 如果生成 token 时发生错误，返回错误响应
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to generate token",
		})
		return
	}

	// 4. 在响应中返回生成的 token
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func ProtectedEndpoint(c *gin.Context) {
	username := c.MustGet("username").(string)
	uuid := c.MustGet("uuid").(string)
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello " + username,
		"uuid":    uuid,
	})
}

package main

import (
	"os"
	"path"

	"github.com/spf13/viper"
	"github.com/x14n/goExperimental/viperExper/common"
)

func InitConfig() {
	workDir, _ := os.Getwd()
	parentDir := path.Dir(workDir)
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(parentDir + "/config")

	err := viper.ReadInConfig()
	if err != nil {
		return
	}
}

func main() {
	InitConfig()
	common.InitDB()
}

// func CollectRoute(r *gin.Engine) *gin.Engine {
// 	// 在这里添加你的路由定义
// 	r.GET("/example", func(c *gin.Context) {
// 		c.JSON(200, gin.H{"message": "Hello, World!"})
// 	})

// 	return r
// }

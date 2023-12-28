package src

import (
	"github.com/gin-gonic/gin"
	"github.com/x14n/goExperimental/gin/server"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/users")

	user.GET("/", server.FindAllUser)
	user.POST("/", server.PostUser)
	user.GET("/:id", server.FindUsreById)
}

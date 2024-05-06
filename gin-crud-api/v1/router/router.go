package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/x14n/goExperimental/gin-crud-api/v1/controller"
)

func NewRouter(tagController *controller.TagController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")
	tagRouter := router.Group("/tag")
	tagRouter.GET("", tagController.FindAll)
	tagRouter.GET("/:tagId", tagController.FindById)
	tagRouter.POST("", tagController.Create)
	tagRouter.PATCH("/:tagId", tagController.Update)
	tagRouter.DELETE("/:tagId", tagController.Delete)

	return service
}

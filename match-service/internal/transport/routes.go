package transport

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, sportHandler *SportHandler) {
	sport := router.Group("/sport")

	{
		sport.GET("", sportHandler.GetList)
		sport.GET("/:id", sportHandler.GetById)
		sport.DELETE("/:id", sportHandler.Delete)
		sport.POST("", sportHandler.Create)
		sport.PATCH("/:id", sportHandler.Update)
	}
}

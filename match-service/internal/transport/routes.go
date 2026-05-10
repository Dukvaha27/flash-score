package transport

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, sportHandler *SportHandler, teamHandler *TeamHandler, playerHandler *PlayerHandler) {
	sport := router.Group("/sport")

	{
		sport.GET("", sportHandler.GetList)
		sport.GET("/:id", sportHandler.GetById)
		sport.DELETE("/:id", sportHandler.Delete)
		sport.POST("", sportHandler.Create)
		sport.PATCH("/:id", sportHandler.Update)
	}

	team := router.Group("/team")

	{
		team.GET("", teamHandler.GetList)
		team.GET("/:id", teamHandler.GetById)
		team.GET("/sport/:id", teamHandler.GetBySport)
		team.POST("", teamHandler.Create)
		team.PATCH("/:id", teamHandler.Update)
		team.DELETE("/:id", teamHandler.Delete)
	}

	player := router.Group("/player")

	{
		player.GET("", playerHandler.GetList)
		player.GET("/:id", playerHandler.GetById)
		player.GET("/team/:id", playerHandler.GetByTeam)
		player.POST("", playerHandler.Create)
		player.PATCH("/:id", playerHandler.Update)
		player.DELETE("/:id", playerHandler.Delete)
	}
}

package transport

import (
	"github.com/Dukvaha27/flash-score/user-service/internal/service"
	"github.com/Dukvaha27/flash-score/user-service/internal/transport/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	userService service.UserService,
) {

	authorized := router.Group("")
	authorized.Use(middleware.InternalAuthMiddleware())

	unauthorized := router.Group("")


	userHandler := NewUserHandler(userService)

	userHandler.RegisterRoutes(authorized, unauthorized)
}

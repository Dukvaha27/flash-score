package transport

import (
	"github.com/Dukvaha27/flash-score/user-service/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	userService service.UserService,
) {

	unauthorized := router.Group("")

	userHandler := NewUserHandler(userService)

	userHandler.RegisterRoutes(unauthorized)
}

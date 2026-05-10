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

	// создать группу защищенную

	userHandler := NewUserHandler(userService)

	userHandler.RegisterRoutes(unauthorized,authorized) // прокинуть сюда защищенную тоже
	// потом внутри от защищенной группы сделать все роуты, где нужен user_id из контекста
}

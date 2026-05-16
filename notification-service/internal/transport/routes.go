package transport

import (
	"github.com/Dukvaha27/flash-score/notification-service/internal/transport/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, subscriptionHandler *SubscriptionHandler) {
	subscription := router.Group("/subscription")
	subscription.Use(middleware.InternalAuthMiddleware())
	{
		subscription.POST("", subscriptionHandler.Subscribe)
		subscription.DELETE("/:subscriptionID", subscriptionHandler.Unsubscribe)
	}
}

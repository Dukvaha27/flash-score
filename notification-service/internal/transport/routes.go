package transport

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine, subscriptionHandler *SubscriptionHandler) {
	subscription := router.Group("/subscription")

	{
		subscription.POST("", subscriptionHandler.Subscribe)
		subscription.DELETE("/:subscriptionID", subscriptionHandler.Unsubscribe)
	}
}

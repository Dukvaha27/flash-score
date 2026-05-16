package transport

import (
	"github.com/Dukvaha27/flash-score/notification-service/internal/transport/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, notificationHandler *NotificationHandler, subscriptionHandler *SubscriptionHandler) {
	notification := router.Group("/notification")
	notification.Use(middleware.InternalAuthMiddleware())
	{
		notification.GET("/unread", notificationHandler.GetUnreadCount)
		notification.GET("/:notificationID", notificationHandler.GetByID)
		notification.PATCH("/:notificationID", notificationHandler.MarkAsRead)
		notification.DELETE("/:notificationID", notificationHandler.Delete)
		notification.POST("", notificationHandler.Create)
	}

	subscription := router.Group("/subscription")
	subscription.Use(middleware.InternalAuthMiddleware())
	{
		subscription.POST("", subscriptionHandler.Subscribe)
		subscription.DELETE("/:subscriptionID", subscriptionHandler.Unsubscribe)
	}

}

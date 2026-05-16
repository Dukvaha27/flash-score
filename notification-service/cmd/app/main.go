package main

import (
	"log"

	"github.com/Dukvaha27/flash-score/notification-service/internal/config"
	"github.com/Dukvaha27/flash-score/notification-service/internal/models"
	"github.com/Dukvaha27/flash-score/notification-service/internal/repository"
	"github.com/Dukvaha27/flash-score/notification-service/internal/service"
	"github.com/Dukvaha27/flash-score/notification-service/internal/transport"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetUpDatabaseConnection()
	if err := db.AutoMigrate(&models.Notification{}, &models.Subscription{}); err != nil {
		log.Fatalf("не удалось выполнить миграции: %v", err)
	}

	router := gin.Default()

	notificationRepo := repository.NewNotificationRepository(db)
	notificationService := service.NewNotificationService(notificationRepo)
	notificationHandler := transport.NewNotificationHandler(notificationService)

	subscriptionRepo := repository.NewSubscriptionRepository(db)
	subscriptionService := service.NewSubscriptionService(subscriptionRepo)
	subscriptionHandler := transport.NewSubscriptionHandler(subscriptionService)

	transport.RegisterRoutes(router, notificationHandler, subscriptionHandler)

	if err := router.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}

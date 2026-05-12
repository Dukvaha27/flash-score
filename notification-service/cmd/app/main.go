package main

import (
	"log"

	"github.com/Dukvaha27/flash-score/notification-service/internal/config"
	"github.com/Dukvaha27/flash-score/notification-service/internal/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetUpDatabaseConnection()
	if err := db.AutoMigrate(&models.Notification{}); err != nil {
		log.Fatalf("не удалось выполнить миграции: %v", err)
	}

	router := gin.Default()

	if err := router.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}

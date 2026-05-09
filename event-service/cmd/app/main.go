package main

import (
	"log"

	"github.com/Dukvaha27/flash-score/event-service/internal/config"
	"github.com/Dukvaha27/flash-score/event-service/internal/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetUpDatebaseConnection()

	if err := db.AutoMigrate(
		&models.MatchEvent{},
		&models.Commentary{},
		&models.Comment{},
		&models.Reaction{},
	); err != nil {
		log.Fatalf("не удалось выполнить миграции: %v", err)
	}

	// ============ Инициализация ============

	// ============ GIN ============

	router := gin.Default()

	if err := router.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}

}

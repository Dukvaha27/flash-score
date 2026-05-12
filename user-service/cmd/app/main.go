package main

import (
	"log"
	"os"

	"github.com/Dukvaha27/flash-score/user-service/internal/config"
	"github.com/Dukvaha27/flash-score/user-service/internal/models"
	"github.com/Dukvaha27/flash-score/user-service/internal/repository"
	"github.com/Dukvaha27/flash-score/user-service/internal/service"
	"github.com/Dukvaha27/flash-score/user-service/internal/transport"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetUpDatabaseConnection()
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET is not set")
		return
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("не удалось выполнить миграции: %v", err)
	}

	router := gin.Default()
	userRepo := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepo, secret)

	transport.RegisterRoutes(router, userService)

	if err := router.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}

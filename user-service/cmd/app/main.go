package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Dukvaha27/flash-score/notification-service/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func main() {
	db := config.SetUpDatebaseConnection()

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	defer rdb.Close()

	ctx := context.Background()

	pong, err := rdb.Ping(ctx).Result()

	if err != nil {
		log.Fatalf("Не удалось подключиться к Redis:", err)
	}

	fmt.Println("Redis подключён", pong)

	if err := db.AutoMigrate(); err != nil {
		log.Fatalf("не удалось выполнить миграции: %v", err)
	}

	// ============ Инициализация ============

	// ============ GIN ============

	router := gin.Default()

	if err := router.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}

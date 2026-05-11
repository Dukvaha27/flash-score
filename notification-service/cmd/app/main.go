package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Dukvaha27/flash-score/notification-service/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {
	db := config.SetUpDatabaseConnection()
	// пока не использую но вскоре использую, ты не комментируй как ошибку.

	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("NOTIFICATION_REDIS_URL"),
	})

	defer rdb.Close()

	ctx := context.Background()

	pong, err := rdb.Ping(ctx).Result()

	if err != nil {
		log.Fatalf("Не удалось подключиться к Redis: %v", err)
	}

	fmt.Println("Redis подключён", pong)

	if err := db.AutoMigrate(); err != nil {
		log.Fatalf("не удалось выполнить миграции: %v", err)
	}

	router := gin.Default()

	if err := router.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}

package main

import (
	"log"

	"github.com/Dukvaha27/flash-score/event-service/internal/config"
	"github.com/Dukvaha27/flash-score/event-service/internal/models"
	"github.com/Dukvaha27/flash-score/event-service/internal/repositories"
	"github.com/Dukvaha27/flash-score/event-service/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetUpDatabaseConnection()

	if err := db.AutoMigrate(
		&models.MatchEvent{},
		&models.Commentary{},
		&models.Comment{},
		&models.Reaction{},
	); err != nil {
		log.Fatalf("не удалось выполнить миграции: %v", err)
	}

	// ============ Repositories ============

	matchEventRepo := repositories.NewMatchEventRepository(db)
	commentaryRepo := repositories.NewCommentaryRepository(db)
	commentRepo := repositories.NewCommentRepository(db)
	reactionRepo := repositories.NewReactionRepository(db)

	// ============ Services ============
	// MatchEventService и CommentaryService пока не создаём,
	// потому что для них нужен реальный MatchClient.
	// Их нужно будет подключить после реализации internal/clients/match_client.go.

	commentService := services.NewCommentService(
		commentRepo,
		matchEventRepo,
		commentaryRepo,
	)

	reactionService := services.NewReactionService(
		reactionRepo,
		matchEventRepo,
		commentaryRepo,
	)

	timelineService := services.NewTimelineService(
		matchEventRepo,
		commentaryRepo,
	)

	// TODO: pass services to HTTP handlers after handlers are implemented
	_ = commentService
	_ = reactionService
	_ = timelineService

	// ============ Gin ============

	router := gin.Default()

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}

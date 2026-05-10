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

	// ============ Инициализация ============

	matchEventRepo := repositories.NewMatchEventRepository(db)
	commentaryRepo := repositories.NewCommentaryRepository(db)
	commentRepo := repositories.NewCommentRepository(db)
	reactionRepo := repositories.NewReactionRepository(db)

	matchEventService := services.NewMatchEventService(
		matchEventRepo,
		nil,
	)

	commentaryService := services.NewCommentaryService(
		db,
		commentaryRepo,
		nil,
	)

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

	_ = matchEventService
	_ = commentaryService
	_ = commentService
	_ = reactionService
	_ = timelineService

	// ============ GIN ============

	router := gin.Default()

	if err := router.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}

}

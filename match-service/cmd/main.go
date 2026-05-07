package main

import (
	"github.com/Dukvaha27/flash-score/match-service/internal/config"
	"github.com/Dukvaha27/flash-score/match-service/internal/models"
	"github.com/Dukvaha27/flash-score/match-service/internal/repo"
	"github.com/Dukvaha27/flash-score/match-service/internal/service"
	"github.com/Dukvaha27/flash-score/match-service/internal/transport"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := config.SetupDatabase()
	db.AutoMigrate(&models.Sport{}, &models.Team{})

	sportRepo := repo.NewSportRepo(db)
	sportService := service.NewSportService(sportRepo)
	sportHandler := transport.NewSportHandler(sportService)

	teamRepo := repo.NewTeamRepo(db)
	teamService := service.NewTeamService(teamRepo, sportRepo)
	teamHandler := transport.NewTeamHandler(teamService)

	transport.RegisterRoutes(router, sportHandler, teamHandler)
	router.Run(":8000")
}

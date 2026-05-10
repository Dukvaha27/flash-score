package transport

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Dukvaha27/flash-score/match-service/internal/models"
	"github.com/Dukvaha27/flash-score/match-service/internal/service"
	"github.com/gin-gonic/gin"
)

type PlayerHandler struct {
	playerService service.PlayerService
}

func NewPlayerHandler(service service.PlayerService) *PlayerHandler {
	return &PlayerHandler{playerService: service}
}

func (p *PlayerHandler) Create(ctx *gin.Context) {
	var req models.PlayerCreate

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	player, err := p.playerService.Create(req)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"err": service.ErrNotFound.Error()})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": player})
}

func (p *PlayerHandler) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	player, err := p.playerService.GetById(uint(id))

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"err": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": player})
}

func (p *PlayerHandler) GetByTeam(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	list, err := p.playerService.GetByTeam(uint(id))

	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"err": service.ErrNotFound.Error()})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		}

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": list})
}

func (p *PlayerHandler) GetList(ctx *gin.Context) {
	list, err := p.playerService.GetList()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": list})
}

func (p *PlayerHandler) Update(ctx *gin.Context) {
	var player models.PlayerUpdate
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&player); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if err := p.playerService.Update(uint(id), &player); err != nil {
		if errors.Is(err, service.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"err": service.ErrNotFound.Error()})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, "updated")
}

func (p *PlayerHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if err := p.playerService.HardDelete(uint(id)); err != nil {
		if errors.Is(err, service.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"err": service.ErrNotFound.Error()})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, "deleted")
}

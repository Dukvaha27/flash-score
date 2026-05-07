package transport

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Dukvaha27/flash-score/match-service/internal/models"
	"github.com/Dukvaha27/flash-score/match-service/internal/service"
	"github.com/gin-gonic/gin"
)

type TeamHandler struct {
	teamService service.TeamService
}

func NewTeamHandler(service service.TeamService) *TeamHandler {
	return &TeamHandler{teamService: service}
}

func (t *TeamHandler) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	team, err := t.teamService.GetById(uint(id))

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": team})
}

func (t *TeamHandler) GetBySport(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	teams, err := t.teamService.GetBySport(uint(id))

	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"err": service.ErrNotFound.Error()})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": teams})
}

func (t *TeamHandler) GetList(ctx *gin.Context) {
	teams, err := t.teamService.GetList()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": teams})
}

func (t *TeamHandler) Create(ctx *gin.Context) {
	var req models.TeamCreate

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	team, err := t.teamService.Create(req)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"err": service.ErrNotFound.Error()})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": team})
}

func (t *TeamHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	var req models.TeamUpdate

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := t.teamService.Update(uint(id), req); err != nil {
		if errors.Is(err, service.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"err": service.ErrNotFound.Error()})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, "ok")
}

func (t *TeamHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if err := t.teamService.HardDelete(uint(id)); err != nil {
		if errors.Is(err, service.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"err": service.ErrNotFound.Error()})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, "deleted")
}

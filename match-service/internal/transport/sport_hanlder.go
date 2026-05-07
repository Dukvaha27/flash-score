package transport

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Dukvaha27/flash-score/match-service/internal/models"
	"github.com/Dukvaha27/flash-score/match-service/internal/service"
	"github.com/gin-gonic/gin"
)

type SportHandler struct {
	sportService service.SportService
}

func NewSportHandler(service service.SportService) *SportHandler {
	return &SportHandler{sportService: service}
}

func (s *SportHandler) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	sport, err := s.sportService.GetById(uint(id))

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"err": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": sport})

}

func (s *SportHandler) GetList(ctx *gin.Context) {
	sports, err := s.sportService.GetList()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": sports})
}

func (s *SportHandler) Update(ctx *gin.Context) {
	var req models.SportAction
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if err := s.sportService.Update(uint(id), req.Name); err != nil {
		if errors.Is(err, service.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"err": service.ErrNotFound.Error()})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, "ok")
}

func (s *SportHandler) Create(ctx *gin.Context) {
	var req models.SportAction

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	sport, err := s.sportService.Create(req.Name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": sport})
}

func (s *SportHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if err := s.sportService.Delete(uint(id)); err != nil {

		if errors.Is(err, service.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"err": service.ErrNotFound.Error()})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, "ok")

}

package transport

import (
	"errors"
	"net/http"
	"strconv"

	myErrors "github.com/Dukvaha27/flash-score/notification-service/internal/errors"
	"github.com/Dukvaha27/flash-score/notification-service/internal/models"
	"github.com/Dukvaha27/flash-score/notification-service/internal/service"
	"github.com/gin-gonic/gin"
)

type SubscriptionHandler struct {
	service service.SubscriptionService
}

func NewSubscriptionHandler(service service.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{service: service}
}

func (h *SubscriptionHandler) Unsubscribe(c *gin.Context) {
	subscriptionID, err := strconv.Atoi(c.Param("subscriptionID"))
	if err != nil || subscriptionID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": myErrors.ErrInvalidType.Error()})
		return
	}
	userID, ok := getUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": myErrors.ErrUnauthorized.Error()})
		return
	}
	if err := h.service.Unsubscribe(uint(subscriptionID), userID); err != nil {
		if errors.Is(err, myErrors.ErrSubscriptionNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Подписка успешно удалена")
}

func (h *SubscriptionHandler) Subscribe(c *gin.Context) {
	var req models.SubscriptionCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Не корректный формат запроса",
		})
		return
	}
	userID, ok := getUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": myErrors.ErrUnauthorized.Error()})
		return
	}

	if err := h.service.Subscribe(req, userID); err != nil {
		if errors.Is(err, myErrors.ErrSubscriptionAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		} else if errors.Is(err, myErrors.ErrSubscriptionTargetRequired) || errors.Is(err, myErrors.ErrTeamOrSportRequired) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, "Подписка создана")
}

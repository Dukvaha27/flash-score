package transport

import (
	"errors"
	"net/http"
	"strconv"

	MyErrors "github.com/Dukvaha27/flash-score/notification-service/internal/errors"
	"github.com/Dukvaha27/flash-score/notification-service/internal/models"
	"github.com/Dukvaha27/flash-score/notification-service/internal/service"
	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	notificationService service.NotificationService
}

func NewNotificationHandler(service service.NotificationService) *NotificationHandler {
	return &NotificationHandler{notificationService: service}
}

func (h *NotificationHandler) Delete(c *gin.Context) {
	userID, ok := getUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": MyErrors.ErrUnauthorized.Error()})
		return
	}
	notificationID, err := strconv.Atoi(c.Param("notificationID"))
	if err != nil || notificationID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": MyErrors.ErrInvalidType.Error()})
		return
	}
	if err := h.notificationService.Delete(uint(notificationID), userID); err != nil {
		if errors.Is(err, MyErrors.ErrNotificationNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": MyErrors.ErrNotificationNotFound.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Уведомление успешно удалено")
}

func (h *NotificationHandler) Create(c *gin.Context) {
	var req models.NotificationCreate
	userID, ok := getUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": MyErrors.ErrUnauthorized.Error()})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Неправильный формат данных"})
		return
	}

	notification, err := h.notificationService.Create(req, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, notification)
}

func (h *NotificationHandler) GetByID(c *gin.Context) {
	userID, ok := getUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": MyErrors.ErrUnauthorized.Error()})
		return
	}
	notificationID, err := strconv.Atoi(c.Param("notificationID"))
	if err != nil || notificationID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": MyErrors.ErrInvalidType.Error()})
		return
	}
	notification, err := h.notificationService.GetByID(uint(notificationID), userID)
	if err != nil {
		if errors.Is(err, MyErrors.ErrNotificationNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": MyErrors.ErrNotificationNotFound.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notification)
}

func (h *NotificationHandler) GetUnreadCount(c *gin.Context) {
	userID, ok := getUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": MyErrors.ErrUnauthorized.Error()})
		return
	}
	count, err := h.notificationService.GetUnreadCount(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, count)
}

func (h *NotificationHandler) MarkAsRead(c *gin.Context) {
	userID, ok := getUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": MyErrors.ErrUnauthorized.Error()})
		return
	}
	notificationID, err := strconv.Atoi(c.Param("notificationID"))
	if err != nil || notificationID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": MyErrors.ErrInvalidType.Error()})
		return
	}
	if err := h.notificationService.MarkAsRead(uint(notificationID), userID); err != nil {
		if errors.Is(err, MyErrors.ErrNotificationNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": MyErrors.ErrNotificationNotFound.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Сообщение прочитано")

}

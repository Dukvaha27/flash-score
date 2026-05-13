package transport

import (
	"errors"
	"net/http"

	"github.com/Dukvaha27/flash-score/user-service/internal/models"
	"github.com/Dukvaha27/flash-score/user-service/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return UserHandler{service: userService}
}

func (h *UserHandler) RegisterRoutes(authorized *gin.RouterGroup, unauthorized *gin.RouterGroup) {
	unauthorized.POST("/register", h.Register)
	unauthorized.POST("/login", h.Login)

	users := authorized.Group("/users")
	{
		users.GET("/me", h.GetByID)
		users.PATCH("/me", h.Update)
		users.DELETE("/me", h.Delete)
	}
}

func (h *UserHandler) GetByID(c *gin.Context) {
	userID, ok := h.getUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	user, err := h.service.GetByID(userID)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Delete(c *gin.Context) {
	userID, ok := h.getUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return

	}
	err := h.service.Delete(userID)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) { // Проверяем специальную ошибку
			c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении"})
		}
		return
	}
	c.JSON(http.StatusOK, "Пользователь успешно удален")
}

func (h *UserHandler) Update(c *gin.Context) {
	var req models.UserUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "Некорректный формат запроса",
			"details": err.Error(),
		})
		return
	}
	userID, ok := h.getUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не авторизован"})
		return
	}
	if err := h.service.Update(userID, req); err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "ошибка при обновлении данных",
		})
		return
	}
	c.JSON(http.StatusOK, "пользователь успешно обновлен")
}

func (h *UserHandler) Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "Некорректный формат запроса",
			"details": err.Error(),
		})
		return
	}

	user, err := h.service.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Ошибка регистрации пользователя",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "Некорректный формат запроса",
			"details": err.Error(),
		})
		return
	}

	token, err := h.service.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Ошибка логина пользователя",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.LoginResponse{Token: token})
}

func (h *UserHandler) getUserIDFromContext(c *gin.Context) (uint, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0, false
	}
	id, ok := userID.(uint)
	return id, ok
}

package transport

import (
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
	rawUserID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Не удалось достать  user_id",
		})
		return
	}
	userID, ok := rawUserID.(uint)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "неправильный тип userID",
		})
		return
	}
	user, err := h.service.GetByID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ошибка при нахождении пользователя",
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Delete(c *gin.Context) {
	rawUserID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Не удалось достать  user_id",
		})
		return
	}
	userID, ok := rawUserID.(uint)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "неправильный тип userID",
		})
		return
	}
	err := h.service.Delete(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ошибка при удалении данных",
		})
		return
	}
	c.JSON(http.StatusOK, "Пользователь успешно удален")
}

func (h *UserHandler) Update(c *gin.Context) {
	var req models.UserUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Некорректный формат запроса",
			"details": err.Error(),
		})
		return
	}
	rawUserID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Не удалось достать  user_id",
		})
		return
	}
	userID, ok := rawUserID.(uint)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "неправильный тип userID",
		})
		return
	}
	if err := h.service.Update(userID, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ошибка при обновлении данных",
		})
		return
	}
	c.JSON(http.StatusOK, "пользователь успешно обновлен")
}

func (h *UserHandler) Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
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
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Некорректный формат запроса",
			"details": err.Error(),
		})
		return
	}

	token, err := h.service.Login(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Ошибка логина пользователя",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.LoginResponse{Token: token})
}

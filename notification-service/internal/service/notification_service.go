package service

import (
	"errors"

	myErrors "github.com/Dukvaha27/flash-score/notification-service/internal/errors"

	"github.com/Dukvaha27/flash-score/notification-service/internal/models"
	"github.com/Dukvaha27/flash-score/notification-service/internal/repository"
	"gorm.io/gorm"
)

type NotificationService interface {
	GetUnreadCount(userID uint) (int64, error)
	MarkAsRead(notificationID, userID uint) error
	GetByID(notificationID, userID uint) (models.Notification, error)
	Delete(notificationID, userID uint) error
	Create(req models.NotificationCreate, userID uint) (*models.Notification, error)
}

type notificationService struct {
	notificationRepo repository.NotificationRepository
}

func NewNotificationService(notificationRepo repository.NotificationRepository) NotificationService {
	return &notificationService{notificationRepo: notificationRepo}
}

func (s *notificationService) GetUnreadCount(userID uint) (int64, error) {
	return s.notificationRepo.GetUnreadCount(userID)
}

func (s *notificationService) Create(req models.NotificationCreate, userID uint) (*models.Notification, error) {
	notification := &models.Notification{Message: req.Message, UserID: userID}
	err := s.notificationRepo.Create(notification)
	if err != nil {
		return nil, err
	}
	return notification, nil
}

func (s *notificationService) MarkAsRead(notificationID, userID uint) error {
	return s.notificationRepo.MarkAsRead(notificationID, userID)
}

func (s *notificationService) GetByID(notificationID, userID uint) (models.Notification, error) {
	notification, err := s.notificationRepo.GetByID(notificationID, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Notification{}, myErrors.ErrNotificationNotFound
		}
		return models.Notification{}, err
	}

	return notification, nil
}

func (s *notificationService) Delete(notificationID, userID uint) error {
	return s.notificationRepo.Delete(notificationID, userID)
}

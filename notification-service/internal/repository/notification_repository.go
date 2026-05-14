package repository

import (
	 "github.com/Dukvaha27/flash-score/notification-service/internal/errors"
	"github.com/Dukvaha27/flash-score/notification-service/internal/models"
	"gorm.io/gorm"
)


type NotificationRepository interface {
	GetUnreadCount(userID uint) (int64, error)
	Create(notification *models.Notification) error
	MarkAsRead(notificationID, userID uint) error
	GetByID(notificationID, userID uint) (models.Notification, error)
	Delete(notificationID, userID uint) error
}

type gormNotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &gormNotificationRepository{db: db}
}

func (r *gormNotificationRepository) GetUnreadCount(userID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Notification{}).Where("user_id = ? and is_read = ?", userID, false).Count(&count).Error
	return count, err
}

func (r *gormNotificationRepository) Create(notification *models.Notification) error {
	return r.db.Create(&notification).Error
}

func (r *gormNotificationRepository) MarkAsRead(notificationID, userID uint) error {
	result := r.db.Model(&models.Notification{}).Where("id = ? and user_id = ?", notificationID, userID).Update("is_read", true)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.ErrNotificationNotFound
	}

	return nil
}

func (r *gormNotificationRepository) GetByID(notificationID, userID uint) (models.Notification, error) {
	var notification models.Notification
	err := r.db.Where("id = ? and user_id = ?", notificationID, userID).First(&notification).Error
	return notification, err
}

func (r *gormNotificationRepository) Delete(notificationID, userID uint) error {
	result := r.db.Where("id = ? and user_id = ?", notificationID, userID).Delete(&models.Notification{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.ErrNotificationNotFound
	}
	return nil
}

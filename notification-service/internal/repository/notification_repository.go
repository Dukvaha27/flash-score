package repository

import (
	"github.com/Dukvaha27/flash-score/notification-service/internal/models"
	"gorm.io/gorm"
)

type NotificationRepository interface {
	Create(notification models.Notification) error
	UpdateStatus(notificationID uint, status bool) error
	GetByID(notificationID uint) (models.Notification, error)
}

type gormNotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &gormNotificationRepository{db: db}
}

func (r *gormNotificationRepository) Create(notification models.Notification) error {
	return r.db.Model(&models.Notification{}).Create(&notification).Error
}

func (r *gormNotificationRepository) UpdateStatus(notificationID uint, status bool) error {
	return r.db.Model(&models.Notification{}).Where("id = ?", notificationID).Update("is_read", status).Error
}

func (r *gormNotificationRepository) GetByID(notificationID uint) (models.Notification, error) {
	var notification models.Notification
	err := r.db.Model(models.Notification{}).Where("id = ?", notificationID).First(&notification).Error
	return notification, err
}

func (r *gormNotificationRepository) Delete(notificationID uint) error {
	return r.db.Model(models.Notification{}).Where("id = ?", notificationID).Delete(&models.Notification{}).Error
}

package repository

import (
	"github.com/Dukvaha27/flash-score/notification-service/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SubscriptionRepository interface {
	Create(subscription models.Subscription) error
	Delete(subscriptionID uint) error
	GetSubscriberIDsByTeam(teamID uint) ([]uint, error)
	GetSubscriberIDsBySport(sportID uint) ([]uint, error)
}

type gormSubscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) SubscriptionRepository {
	return &gormSubscriptionRepository{db: db}
}

func (r *gormSubscriptionRepository) Create(subscription models.Subscription) error {
	return r.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&subscription).Error
}

func (r *gormSubscriptionRepository) Delete(subscriptionID uint) error {
	return r.db.Delete(&models.Subscription{}, subscriptionID).Error
}

func (r *gormSubscriptionRepository) GetSubscriberIDsByTeam(teamID uint) ([]uint, error) {
	var userIDs []uint
	err := r.db.Model(&models.Subscription{}).Where("team_id = ?", teamID).Pluck("user_id", &userIDs).Error
	return userIDs, err
}

func (r *gormSubscriptionRepository) GetSubscriberIDsBySport(sportID uint) ([]uint, error) {
	var userIDs []uint
	err := r.db.Model(&models.Subscription{}).Where("sport_id = ?", sportID).Pluck("user_id", &userIDs).Error
	return userIDs, err
}

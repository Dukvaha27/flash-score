package repository

import (
	"github.com/Dukvaha27/flash-score/notification-service/internal/errors"
	"github.com/Dukvaha27/flash-score/notification-service/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SubscriptionRepository interface {
	Subscribe(subscription models.Subscription) error
	Unsubscribe(subscriptionID, userID uint) error
	GetSubscriberIDsByTeam(teamID uint) ([]uint, error)
	GetSubscriberIDsBySport(sportID uint) ([]uint, error)
}

type gormSubscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) SubscriptionRepository {
	return &gormSubscriptionRepository{db: db}
}

func (r *gormSubscriptionRepository) Subscribe(subscription models.Subscription) error {
	result := r.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&subscription)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.ErrSubscriptionAlreadyExists
	}
	return nil
}

func (r *gormSubscriptionRepository) Unsubscribe(subscriptionID, userID uint) error {
	result := r.db.Where("id = ? and user_id = ?", subscriptionID, userID).Delete(&models.Subscription{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.ErrSubscriptionNotFound
	}
	return nil
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

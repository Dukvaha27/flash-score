package service

import (
	"github.com/Dukvaha27/flash-score/notification-service/internal/models"
	"github.com/Dukvaha27/flash-score/notification-service/internal/repository"
)

type SubscriptionService interface {
	Subscribe(req models.SubscriptionCreate, userID uint) error
	Unsubscribe(subscriptionID, userID uint) error
}

type subscriptionService struct {
	subscriptionRepo repository.SubscriptionRepository
}

func NewSubscriptionService(subscriptionRepo repository.SubscriptionRepository) SubscriptionService {
	return &subscriptionService{subscriptionRepo: subscriptionRepo}
}

func (s *subscriptionService) Subscribe(req models.SubscriptionCreate, userID uint) error {
	if err := req.Validate(); err != nil {
		return err
	}
	subscription := &models.Subscription{
		UserID:  userID,
		TeamID:  req.TeamID,
		SportID: req.SportID,
	}
	return s.subscriptionRepo.Subscribe(*subscription)
}

func (s *subscriptionService) Unsubscribe(subscriptionID, userID uint) error {
	return s.subscriptionRepo.Unsubscribe(subscriptionID, userID)
}

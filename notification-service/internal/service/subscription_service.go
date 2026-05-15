package service

import (
	"github.com/Dukvaha27/flash-score/notification-service/internal/models"
	"github.com/Dukvaha27/flash-score/notification-service/internal/repository"
)

type SubscriptionService interface {
	Subscribe(req models.SubscriptionCreate, userID uint) error
	Unsubscribe(subscriptionID, userID uint) error
	GetSubscriberIDsByTeam(teamID uint) ([]uint, error)
	GetSubscriberIDsBySport(sportID uint) ([]uint, error)
}

type subscriptionService struct {
	subscriptionRepo repository.SubscriptionRepository
}

func NewSubscriptionService(subscriptionRepo repository.SubscriptionRepository) SubscriptionService {
	return &subscriptionService{subscriptionRepo: subscriptionRepo}
}

func (s *subscriptionService) GetSubscriberIDsBySport(sportID uint) ([]uint, error) {
	userIDs, err := s.subscriptionRepo.GetSubscriberIDsBySport(sportID)
	if err != nil {
		return nil, err
	}

	return userIDs, nil
}

func (s *subscriptionService) GetSubscriberIDsByTeam(teamID uint) ([]uint, error) {
	userIDs, err := s.subscriptionRepo.GetSubscriberIDsByTeam(teamID)
	if err != nil {
		return nil, err
	}

	return userIDs, nil
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

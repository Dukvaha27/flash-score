package services

import (
	"context"

	"github.com/Dukvaha27/flash-score/event-service/internal/dto"
	"github.com/Dukvaha27/flash-score/event-service/internal/models"
	"github.com/Dukvaha27/flash-score/event-service/internal/repositories"
)

type MatchEventService interface {
	Create(ctx context.Context, req dto.CreateEventRequest, userID uint, role string) (*models.MatchEvent, error)
	GetByID(id uint) (*models.MatchEvent, error)
	Update(ctx context.Context, id uint, req dto.UpdateEventRequest, role string) error
	Delete(id uint, role string) error
	ListByMatchID(matchID uint) ([]models.MatchEvent, error)
}

type matchEventService struct {
	eventRepo   repositories.MatchEventRepository
	matchClient MatchClient
}

func NewMatchEventService(
	eventRepo repositories.MatchEventRepository,
	matchClient MatchClient,
) (MatchEventService, error) {
	if matchClient == nil {
		return nil, ErrMatchClientNotConfigured
	}

	return &matchEventService{
		eventRepo:   eventRepo,
		matchClient: matchClient,
	}, nil
}

func (s *matchEventService) Create(ctx context.Context, req dto.CreateEventRequest, userID uint, role string) (*models.MatchEvent, error) {
	if !canManageMatchContent(role) {
		return nil, ErrForbidden
	}

	if !isValidEventType(req.EventType) {
		return nil, ErrInvalidEventType
	}

	match, err := s.matchClient.GetMatchByID(ctx, req.MatchID)
	if err != nil {
		return nil, err
	}

	if match.Status != "live" {
		return nil, ErrMatchNotLive
	}

	if req.EventType == string(models.Goal) {
		if req.TeamID == nil {
			return nil, ErrTeamRequired
		}

		if *req.TeamID != match.HomeTeamID && *req.TeamID != match.AwayTeamID {
			return nil, ErrTeamNotInMatch
		}
	}

	event := models.MatchEvent{
		MatchID:   req.MatchID,
		EventType: models.EventType(req.EventType),
		Minute:    req.Minute,
		TeamID:    req.TeamID,
		PlayerID:  req.PlayerID,
		Text:      req.Text,
		CreatedBy: userID,
	}

	if err = s.eventRepo.Create(&event); err != nil {
		return nil, err
	}

	return &event, nil
}

func (s *matchEventService) GetByID(id uint) (*models.MatchEvent, error) {
	return s.eventRepo.GetByID(id)
}

func (s *matchEventService) Update(ctx context.Context, id uint, req dto.UpdateEventRequest, role string) error {
	if !canManageMatchContent(role) {
		return ErrForbidden
	}

	if !isValidEventType(req.EventType) {
		return ErrInvalidEventType
	}

	existing, err := s.eventRepo.GetByID(id)
	if err != nil {
		return err
	}

	match, err := s.matchClient.GetMatchByID(ctx, existing.MatchID)
	if err != nil {
		return err
	}

	if match.Status != "live" {
		return ErrMatchNotLive
	}

	if req.EventType == string(models.Goal) {
		if req.TeamID == nil {
			return ErrTeamRequired
		}

		if *req.TeamID != match.HomeTeamID && *req.TeamID != match.AwayTeamID {
			return ErrTeamNotInMatch
		}
	}

	event := &models.MatchEvent{
		EventType: models.EventType(req.EventType),
		Minute:    req.Minute,
		TeamID:    req.TeamID,
		PlayerID:  req.PlayerID,
		Text:      req.Text,
	}

	return s.eventRepo.Update(id, event)
}

func (s *matchEventService) Delete(id uint, role string) error {
	if !canManageMatchContent(role) {
		return ErrForbidden
	}

	return s.eventRepo.Delete(id)
}

func (s *matchEventService) ListByMatchID(matchID uint) ([]models.MatchEvent, error) {
	return s.eventRepo.ListByMatchID(matchID)

}

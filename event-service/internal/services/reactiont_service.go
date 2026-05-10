package services

import (
	"github.com/Dukvaha27/flash-score/event-service/internal/dto"
	"github.com/Dukvaha27/flash-score/event-service/internal/models"
	"github.com/Dukvaha27/flash-score/event-service/internal/repositories"
)

type ReactionService interface {
	Upsert(req dto.UpsertReactionRequest, userID uint) (*models.Reaction, error)
	Delete(eventID *uint, commentaryID *uint, userID uint) error

	CountByEventID(eventID uint) ([]repositories.ReactionCount, error)
	CountByCommentaryID(commentaryID uint) ([]repositories.ReactionCount, error)
}

type reactionService struct {
	reactionRepo   repositories.ReactionRepository
	eventRepo      repositories.MatchEventRepository
	commentaryRepo repositories.CommentaryRepository
}

func NewReactionService(
	reactionRepo repositories.ReactionRepository,
	eventRepo repositories.MatchEventRepository,
	commentaryRepo repositories.CommentaryRepository,
) ReactionService {
	return &reactionService{
		reactionRepo:   reactionRepo,
		eventRepo:      eventRepo,
		commentaryRepo: commentaryRepo,
	}
}

func (s *reactionService) Upsert(req dto.UpsertReactionRequest, userID uint) (*models.Reaction, error) {
	if !hasExactlyOneTarget(req.EventID, req.CommentaryID) {
		return nil, ErrInvalidTarget
	}

	if !isValidReactionType(req.Type) {
		return nil, ErrInvalidReaction
	}

	if req.EventID != nil {
		if _, err := s.eventRepo.GetByID(*req.EventID); err != nil {
			return nil, err
		}
	}

	if req.CommentaryID != nil {
		if _, err := s.commentaryRepo.GetByID(*req.CommentaryID); err != nil {
			return nil, err
		}
	}

	reaction := &models.Reaction{
		UserID:        userID,
		EventID:       req.EventID,
		CommentaryID:  req.CommentaryID,
		ReactionsType: models.ReactionsType(req.Type),
	}

	return s.reactionRepo.Upsert(reaction)
}

func (s *reactionService) Delete(eventID *uint, commentaryID *uint, userID uint) error {
	if !hasExactlyOneTarget(eventID, commentaryID) {
		return ErrInvalidTarget
	}

	return s.reactionRepo.DeleteByTarget(userID, eventID, commentaryID)
}

func (s *reactionService) CountByEventID(eventID uint) ([]repositories.ReactionCount, error) {
	return s.reactionRepo.CountByEventID(eventID)
}

func (s *reactionService) CountByCommentaryID(commentaryID uint) ([]repositories.ReactionCount, error) {
	return s.reactionRepo.CountByCommentaryID(commentaryID)
}

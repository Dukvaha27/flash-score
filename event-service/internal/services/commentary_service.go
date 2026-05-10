package services

import (
	"context"

	"github.com/Dukvaha27/flash-score/event-service/internal/dto"
	"github.com/Dukvaha27/flash-score/event-service/internal/models"
	"github.com/Dukvaha27/flash-score/event-service/internal/repositories"
	"gorm.io/gorm"
)

type CommentaryService interface {
	Create(ctx context.Context, req dto.CreateCommentaryRequest, userID uint, role string) (*models.Commentary, error)
	GetByID(id uint) (*models.Commentary, error)
	Update(ctx context.Context, id uint, req dto.UpdateCommentaryRequest, role string) error
	Delete(id uint, role string) error

	ListByMatchID(matchID uint) ([]models.Commentary, error)

	Pin(id uint, role string) error
	Unpin(id uint, role string) error
}

type commentaryService struct {
	db             *gorm.DB
	commentaryRepo repositories.CommentaryRepository
	matchClient    MatchClient
}

func NewCommentaryService(
	db *gorm.DB,
	commentaryRepo repositories.CommentaryRepository,
	matchClient MatchClient,
) (CommentaryService, error) {
	if matchClient == nil {
		return nil, ErrMatchClientNotConfigured
	}

	return &commentaryService{
		db:             db,
		commentaryRepo: commentaryRepo,
		matchClient:    matchClient,
	}, nil
}

func (s *commentaryService) Create(
	ctx context.Context,
	req dto.CreateCommentaryRequest,
	userID uint,
	role string,
) (*models.Commentary, error) {
	if !canManageMatchContent(role) {
		return nil, ErrForbidden
	}

	match, err := s.matchClient.GetMatchByID(ctx, req.MatchID)
	if err != nil {
		return nil, err
	}

	if match.Status != "live" {
		return nil, ErrMatchNotLive
	}

	commentary := &models.Commentary{
		MatchID:   req.MatchID,
		Minute:    req.Minute,
		Text:      req.Text,
		IsPinned:  false,
		CreatedBy: userID,
	}

	if !req.IsPinned {
		if err := s.commentaryRepo.Create(commentary); err != nil {
			return nil, err
		}

		return commentary, nil
	}

	err = s.db.Transaction(func(tx *gorm.DB) error {
		repo := s.commentaryRepo.WithDB(tx)

		if err := repo.Create(commentary); err != nil {
			return err
		}

		if err := repo.UnpinAllByMatchID(commentary.MatchID); err != nil {
			return err
		}

		return repo.SetPinned(commentary.ID, true)
	})

	if err != nil {
		return nil, err
	}

	commentary.IsPinned = true

	return commentary, nil
}

func (s *commentaryService) GetByID(id uint) (*models.Commentary, error) {
	return s.commentaryRepo.GetByID(id)
}

func (s *commentaryService) Update(
	ctx context.Context,
	id uint,
	req dto.UpdateCommentaryRequest,
	role string,
) error {
	if !canManageMatchContent(role) {
		return ErrForbidden
	}

	existing, err := s.commentaryRepo.GetByID(id)
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

	commentary := &models.Commentary{
		Minute: req.Minute,
		Text:   req.Text,
	}

	return s.commentaryRepo.Update(id, commentary)
}

func (s *commentaryService) Delete(id uint, role string) error {
	if !canManageMatchContent(role) {
		return ErrForbidden
	}

	return s.commentaryRepo.Delete(id)
}

func (s *commentaryService) ListByMatchID(matchID uint) ([]models.Commentary, error) {
	return s.commentaryRepo.ListByMatchID(matchID)
}

func (s *commentaryService) Pin(id uint, role string) error {
	if !canManageMatchContent(role) {
		return ErrForbidden
	}

	return s.db.Transaction(func(tx *gorm.DB) error {
		repo := s.commentaryRepo.WithDB(tx)

		commentary, err := repo.GetByID(id)
		if err != nil {
			return err
		}

		if err := repo.UnpinAllByMatchID(commentary.MatchID); err != nil {
			return err
		}

		return repo.SetPinned(id, true)
	})
}

func (s *commentaryService) Unpin(id uint, role string) error {
	if !canManageMatchContent(role) {
		return ErrForbidden
	}

	return s.commentaryRepo.SetPinned(id, false)
}

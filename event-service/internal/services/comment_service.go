package services

import (
	"github.com/Dukvaha27/flash-score/event-service/internal/dto"
	"github.com/Dukvaha27/flash-score/event-service/internal/models"
	"github.com/Dukvaha27/flash-score/event-service/internal/repositories"
)

type CommentService interface {
	Create(req dto.CreateCommentRequest, userID uint) (*models.Comment, error)
	GetByID(id uint) (*models.Comment, error)
	Update(id uint, req dto.UpdateCommentRequest, userID uint) error
	Delete(id uint, userID uint, role string) error

	ListByEventID(eventID uint, limit, offset int) ([]models.Comment, int64, error)
	ListByCommentaryID(commentaryID uint, limit, offset int) ([]models.Comment, int64, error)
}

type commentService struct {
	commentRepo    repositories.CommentRepository
	eventRepo      repositories.MatchEventRepository
	commentaryRepo repositories.CommentaryRepository
}

func NewCommentService(
	commentRepo repositories.CommentRepository,
	eventRepo repositories.MatchEventRepository,
	commentaryRepo repositories.CommentaryRepository,
) CommentService {
	return &commentService{
		commentRepo:    commentRepo,
		eventRepo:      eventRepo,
		commentaryRepo: commentaryRepo,
	}
}

func (s *commentService) Create(req dto.CreateCommentRequest, userID uint) (*models.Comment, error) {
	if !hasExactlyOneTarget(req.EventID, req.CommentaryID) {
		return nil, ErrInvalidTarget
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

	comment := &models.Comment{
		UserID:       userID,
		EventID:      req.EventID,
		CommentaryID: req.CommentaryID,
		Text:         req.Text,
	}

	if err := s.commentRepo.Create(comment); err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *commentService) GetByID(id uint) (*models.Comment, error) {
	return s.commentRepo.GetByID(id)
}

func (s *commentService) Update(id uint, req dto.UpdateCommentRequest, userID uint) error {
	comment, err := s.commentRepo.GetByID(id)
	if err != nil {
		return err
	}

	if comment.UserID != userID {
		return ErrForbidden
	}

	updatedComment := &models.Comment{
		Text: req.Text,
	}

	return s.commentRepo.Update(id, updatedComment)
}

func (s *commentService) Delete(id uint, userID uint, role string) error {
	comment, err := s.commentRepo.GetByID(id)
	if err != nil {
		return err
	}

	if !canDeleteUserContent(comment.UserID, userID, role) {
		return ErrForbidden
	}

	return s.commentRepo.Delete(id)
}

func (s *commentService) ListByEventID(eventID uint, limit, offset int) ([]models.Comment, int64, error) {
	return s.commentRepo.ListByEventID(eventID, limit, offset)
}

func (s *commentService) ListByCommentaryID(commentaryID uint, limit, offset int) ([]models.Comment, int64, error) {
	return s.commentRepo.ListByCommentaryID(commentaryID, limit, offset)
}

package repositories

import (
	"github.com/Dukvaha27/flash-score/event-service/internal/models"
	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment *models.Comment) error
	GetByID(id uint) (*models.Comment, error)
	Update(id uint, comment *models.Comment) error
	Delete(id uint) error

	ListByEventID(eventID uint, limit, offset int) ([]models.Comment, int64, error)
	ListByCommentaryID(commentaryID uint, limit, offset int) ([]models.Comment, int64, error)

	WithDB(db *gorm.DB) CommentRepository
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (r *commentRepository) WithDB(db *gorm.DB) CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (r *commentRepository) Create(comment *models.Comment) error {
	return r.db.Create(comment).Error
}

func (r *commentRepository) GetByID(id uint) (*models.Comment, error) {
	var comment models.Comment

	if err := r.db.First(&comment, id).Error; err != nil {
		return nil, err
	}

	return &comment, nil
}

func (r *commentRepository) Update(id uint, comment *models.Comment) error {
	result := r.db.Model(&models.Comment{}).
		Where("id = ?", id).
		Update("text", comment.Text)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *commentRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Comment{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *commentRepository) ListByEventID(eventID uint, limit, offset int) ([]models.Comment, int64, error) {
	limit, offset = normalizeLimitOffset(limit, offset)

	var comments []models.Comment
	var total int64

	query := r.db.Model(&models.Comment{}).
		Where("event_id = ?", eventID).
		Where("commentary_id IS NULL")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.
		Order("created_at ASC").
		Limit(limit).
		Offset(offset).
		Find(&comments).Error

	return comments, total, err
}

func (r *commentRepository) ListByCommentaryID(commentaryID uint, limit, offset int) ([]models.Comment, int64, error) {
	limit, offset = normalizeLimitOffset(limit, offset)

	var comments []models.Comment
	var total int64

	query := r.db.Model(&models.Comment{}).
		Where("commentary_id = ?", commentaryID).
		Where("event_id IS NULL")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.
		Order("created_at ASC").
		Limit(limit).
		Offset(offset).
		Find(&comments).Error

	return comments, total, err
}

func normalizeLimitOffset(limit, offset int) (int, int) {
	if limit <= 0 {
		limit = 20
	}

	if limit > 100 {
		limit = 100
	}

	if offset < 0 {
		offset = 0
	}

	return limit, offset
}

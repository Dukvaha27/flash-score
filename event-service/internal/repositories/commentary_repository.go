package repositories

import (
	"github.com/Dukvaha27/flash-score/event-service/internal/models"
	"gorm.io/gorm"
)

type CommentaryRepository interface {
	Create(commentary *models.Commentary) error
	GetByID(id uint) (*models.Commentary, error)
	Update(id uint, commentary *models.Commentary) error
	Delete(id uint) error

	ListByMatchID(matchID uint) ([]models.Commentary, error)

	Pin(id uint) error
	Unpin(id uint) error

	WithDB(db *gorm.DB) CommentaryRepository
}

type commentaryRepository struct {
	db *gorm.DB
}

func NewCommentaryRepository(db *gorm.DB) CommentaryRepository {
	return &commentaryRepository{
		db: db,
	}
}

func (r *commentaryRepository) WithDB(db *gorm.DB) CommentaryRepository {
	return &commentaryRepository{
		db: db,
	}
}

func (r *commentaryRepository) Create(commentary *models.Commentary) error {
	return r.db.Create(commentary).Error
}

func (r *commentaryRepository) GetByID(id uint) (*models.Commentary, error) {
	var commentary models.Commentary

	if err := r.db.First(&commentary, id).Error; err != nil {
		return nil, err
	}

	return &commentary, nil
}

func (r *commentaryRepository) Update(id uint, commentary *models.Commentary) error {
	result := r.db.Model(&models.Commentary{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"minute": commentary.Minute,
			"text":   commentary.Text,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *commentaryRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Commentary{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *commentaryRepository) ListByMatchID(matchID uint) ([]models.Commentary, error) {
	var commentaries []models.Commentary

	err := r.db.
		Where("match_id = ?", matchID).
		Order("minute ASC").
		Order("created_at ASC").
		Find(&commentaries).Error

	return commentaries, err
}

func (r *commentaryRepository) Pin(id uint) error {
	var commentary models.Commentary

	if err := r.db.First(&commentary, id).Error; err != nil {
		return err
	}

	if err := r.db.Model(&models.Commentary{}).
		Where("match_id = ?", commentary.MatchID).
		Update("is_pinned", false).Error; err != nil {
		return err
	}

	result := r.db.Model(&models.Commentary{}).
		Where("id = ?", id).
		Update("is_pinned", true)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *commentaryRepository) Unpin(id uint) error {
	result := r.db.Model(&models.Commentary{}).
		Where("id = ?", id).
		Update("is_pinned", false)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

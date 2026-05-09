package repositories

import (
	"github.com/Dukvaha27/flash-score/event-service/internal/models"
	"gorm.io/gorm"
)

type MatchEventRepository interface {
	Create(event *models.MatchEvent) error
	GetByID(id uint) (*models.MatchEvent, error)
	Update(id uint, event *models.MatchEvent) error
	Delete(id uint) error

	ListByMatchID(matchID uint) ([]models.MatchEvent, error)

	WithDB(db *gorm.DB) MatchEventRepository
}

type matchEventRepository struct {
	db *gorm.DB
}

func NewMatchEventRepository(db *gorm.DB) MatchEventRepository {
	return &matchEventRepository{
		db: db,
	}
}

func (r *matchEventRepository) WithDB(db *gorm.DB) MatchEventRepository {
	return &matchEventRepository{
		db: db,
	}
}

func (r *matchEventRepository) Create(event *models.MatchEvent) error {
	return r.db.Create(event).Error
}

func (r *matchEventRepository) GetByID(id uint) (*models.MatchEvent, error) {
	var event models.MatchEvent

	if err := r.db.First(&event, id).Error; err != nil {
		return nil, err
	}

	return &event, nil
}

func (r *matchEventRepository) Update(id uint, event *models.MatchEvent) error {
	result := r.db.Model(&models.MatchEvent{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"event_type": event.EventType,
			"minute":     event.Minute,
			"team_id":    event.TeamID,
			"player_id":  event.PlayerID,
			"text":       event.Text,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *matchEventRepository) Delete(id uint) error {
	result := r.db.Delete(&models.MatchEvent{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *matchEventRepository) ListByMatchID(matchID uint) ([]models.MatchEvent, error) {
	var events []models.MatchEvent

	err := r.db.
		Where("match_id = ?", matchID).
		Order("minute ASC").
		Order("created_at ASC").
		Find(&events).Error

	return events, err
}

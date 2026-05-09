package repositories

import (
	"errors"

	"github.com/Dukvaha27/flash-score/event-service/internal/models"
	"gorm.io/gorm"
)

type ReactionCount struct {
	Type  string `json:"type"`
	Count int64  `json:"count"`
}

type ReactionRepository interface {
	Upsert(reaction *models.Reaction) (*models.Reaction, error)
	DeleteByTarget(userID uint, eventID *uint, commentaryID *uint) error

	CountByEventID(eventID uint) ([]ReactionCount, error)
	CountByCommentaryID(commentaryID uint) ([]ReactionCount, error)

	WithDB(db *gorm.DB) ReactionRepository
}

type reactionRepository struct {
	db *gorm.DB
}

func NewReactionRepository(db *gorm.DB) ReactionRepository {
	return &reactionRepository{
		db: db,
	}
}

func (r *reactionRepository) WithDB(db *gorm.DB) ReactionRepository {
	return &reactionRepository{
		db: db,
	}
}

func (r *reactionRepository) Upsert(reaction *models.Reaction) (*models.Reaction, error) {
	if !hasExactlyOneTarget(reaction.EventID, reaction.CommentaryID) {
		return nil, errors.New("reaction must target either event or commentary")
	}

	var existing models.Reaction

	query := r.db.Where("user_id = ?", reaction.UserID)

	if reaction.EventID != nil {
		query = query.
			Where("event_id = ?", *reaction.EventID).
			Where("commentary_id IS NULL")
	} else {
		query = query.
			Where("commentary_id = ?", *reaction.CommentaryID).
			Where("event_id IS NULL")
	}

	err := query.First(&existing).Error

	if err == nil {
		existing.ReactionsType = reaction.ReactionsType

		if err := r.db.Save(&existing).Error; err != nil {
			return nil, err
		}

		return &existing, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if err := r.db.Create(reaction).Error; err != nil {
		return nil, err
	}

	return reaction, nil
}

func (r *reactionRepository) DeleteByTarget(userID uint, eventID *uint, commentaryID *uint) error {
	if !hasExactlyOneTarget(eventID, commentaryID) {
		return errors.New("reaction must target either event or commentary")
	}

	query := r.db.Where("user_id = ?", userID)

	if eventID != nil {
		query = query.
			Where("event_id = ?", *eventID).
			Where("commentary_id IS NULL")
	} else {
		query = query.
			Where("commentary_id = ?", *commentaryID).
			Where("event_id IS NULL")
	}

	result := query.Delete(&models.Reaction{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *reactionRepository) CountByEventID(eventID uint) ([]ReactionCount, error) {
	var counts []ReactionCount

	err := r.db.
		Model(&models.Reaction{}).
		Select("reactions_type AS type, COUNT(*) AS count").
		Where("event_id = ?", eventID).
		Where("commentary_id IS NULL").
		Group("reactions_type").
		Scan(&counts).Error

	return counts, err
}

func (r *reactionRepository) CountByCommentaryID(commentaryID uint) ([]ReactionCount, error) {
	var counts []ReactionCount

	err := r.db.
		Model(&models.Reaction{}).
		Select("reactions_type AS type, COUNT(*) AS count").
		Where("commentary_id = ?", commentaryID).
		Where("event_id IS NULL").
		Group("reactions_type").
		Scan(&counts).Error

	return counts, err
}

func hasExactlyOneTarget(eventID *uint, commentaryID *uint) bool {
	return (eventID != nil && commentaryID == nil) ||
		(eventID == nil && commentaryID != nil)
}

package repo

import (
	"github.com/Dukvaha27/flash-score/match-service/internal/models"
	"gorm.io/gorm"
)

type SportRepo interface {
	Create(sport *models.Sport) error
	Delete(id uint) error
	Update(sport *models.Sport) error
	GetList() (*[]models.Sport, error)
	GetById(id uint) (*models.Sport, error)
}

type gormSportRepo struct {
	db *gorm.DB
}

func NewSportRepo(db *gorm.DB) SportRepo {
	return &gormSportRepo{
		db: db,
	}
}

func (s *gormSportRepo) Create(sport *models.Sport) error {
	return s.db.Create(&sport).Error
}

func (s *gormSportRepo) Delete(id uint) error {
	return s.db.Unscoped().Delete(&models.Sport{}, id).Error
}

func (s *gormSportRepo) Update(sport *models.Sport) error {
	return s.db.Save(&sport).Error
}

func (s *gormSportRepo) GetById(id uint) (*models.Sport, error) {
	sport := models.Sport{}

	if err := s.db.First(&sport, id).Error; err != nil {
		return nil, err
	}

	return &sport, nil
}

func (s *gormSportRepo) GetList() (*[]models.Sport, error) {
	sports := []models.Sport{}

	if err := s.db.Find(&sports).Error; err != nil {
		return nil, err
	}

	return &sports, nil
}

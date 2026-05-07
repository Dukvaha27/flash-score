package repo

import (
	"github.com/Dukvaha27/flash-score/match-service/internal/models"
	"gorm.io/gorm"
)

type TeamRepo interface {
	GetById(id uint) (*models.Team, error)
	GetBySport(sportId uint) ([]models.Team, error)
	GetList() ([]models.Team, error)
	Create(team models.Team) (*models.Team, error)
	Update(team models.Team) error
	Delete(id uint) error
}

type gormTeamRepo struct {
	db *gorm.DB
}

func NewTeamRepo(db *gorm.DB) TeamRepo {
	return &gormTeamRepo{db: db}
}

func (t *gormTeamRepo) GetById(id uint) (*models.Team, error) {
	var team models.Team

	if err := t.db.First(&team, id).Error; err != nil {
		return nil, err
	}

	return &team, nil
}

func (t *gormTeamRepo) GetBySport(sportId uint) ([]models.Team, error) {
	var teams []models.Team

	if err := t.db.Where("sport_id = ?", sportId).Find(&teams).Error; err != nil {
		return nil, err
	}
	return teams, nil
}

func (t *gormTeamRepo) GetList() ([]models.Team, error) {
	var teams []models.Team

	if err := t.db.Find(&teams).Error; err != nil {
		return nil, err
	}
	return teams, nil
}

func (t *gormTeamRepo) Create(team models.Team) (*models.Team, error) {
	if err := t.db.Create(&team).Error; err != nil {
		return nil, err
	}
	return &team, nil
}

func (t *gormTeamRepo) Update(team models.Team) error {
	return t.db.Save(team).Error
}

func (t *gormTeamRepo) Delete(id uint) error {
	return t.db.Unscoped().Delete(models.Team{}, id).Error
}

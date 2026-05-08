package repo

import (
	"github.com/Dukvaha27/flash-score/match-service/internal/models"
	"gorm.io/gorm"
)

type PlayerRepo interface {
	Create(player models.Player) (*models.Player, error)
	GetById(id uint) (*models.Player, error)
	GetList() ([]models.Player, error)
	GetByTeam(id uint) ([]models.Player, error)
	Update(player *models.Player) error
	HardDelete(id uint) error
}

type gormPlayerRepo struct {
	db *gorm.DB
}

func NewPlayerRepo(db *gorm.DB) PlayerRepo {
	return &gormPlayerRepo{db: db}
}

func (p *gormPlayerRepo) Create(player models.Player) (*models.Player, error) {
	if err := p.db.Create(&player).Error; err != nil {
		return nil, err
	}
	return &player, nil
}

func (p *gormPlayerRepo) GetById(id uint) (*models.Player, error) {
	var player models.Player

	if err := p.db.First(&player, id).Error; err != nil {
		return nil, err
	}
	return &player, nil
}

func (p *gormPlayerRepo) GetList() ([]models.Player, error) {
	var players []models.Player

	if err := p.db.Find(&players).Error; err != nil {
		return nil, err
	}
	return players, nil
}

func (p *gormPlayerRepo) GetByTeam(id uint) ([]models.Player, error) {
	var players []models.Player

	if err := p.db.Where("team_id = ?", id).Find(&players).Error; err != nil {
		return nil, err
	}
	return players, nil
}

func (p *gormPlayerRepo) Update(player *models.Player) error {
	return p.db.Save(player).Error
}

func (p *gormPlayerRepo) HardDelete(id uint) error {
	result := p.db.Unscoped().Delete(&models.Player{}, id)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

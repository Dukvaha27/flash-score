package service

import (
	"errors"

	"github.com/Dukvaha27/flash-score/match-service/internal/models"
	"github.com/Dukvaha27/flash-score/match-service/internal/repo"
	"gorm.io/gorm"
)

type PlayerService interface {
	Create(player models.PlayerCreate) (*models.Player, error)
	GetById(id uint) (*models.Player, error)
	GetList() ([]models.Player, error)
	GetByTeam(id uint) ([]models.Player, error)
	Update(id uint, player *models.PlayerUpdate) error
	HardDelete(id uint) error
}

type playerService struct {
	playerRepo repo.PlayerRepo
	teamRepo   repo.TeamRepo
}

func NewPlayerService(playerRepo repo.PlayerRepo, teamRepo repo.TeamRepo) PlayerService {
	return &playerService{playerRepo: playerRepo, teamRepo: teamRepo}
}

func (p *playerService) Create(req models.PlayerCreate) (*models.Player, error) {
	if _, err := p.teamRepo.GetById(req.TeamID); err != nil {
		return nil, err
	}
	player := models.Player{
		Name:     req.Name,
		Position: req.Position,
		Number:   req.Number,
		TeamID:   req.TeamID,
	}

	return p.playerRepo.Create(player)
}

func (p *playerService) GetById(id uint) (*models.Player, error) {
	return p.playerRepo.GetById(id)
}

func (p *playerService) GetList() ([]models.Player, error) {
	return p.playerRepo.GetList()
}

func (p *playerService) GetByTeam(id uint) ([]models.Player, error) {
	if _, err := p.teamRepo.GetById(id); err != nil {
		return nil, ErrNotFound
	}
	return p.playerRepo.GetByTeam(id)
}

func (p *playerService) Update(id uint, req *models.PlayerUpdate) error {

	player, err := p.GetById(id)
	if err != nil {
		return ErrNotFound
	}

	if req.Name != nil {
		player.Name = *req.Name
	}
	if req.Number != nil {
		player.Number = *req.Number
	}
	if req.Position != nil {
		player.Position = *req.Position
	}
	if req.TeamID != nil {
		if _, err := p.teamRepo.GetById(*req.TeamID); err != nil {
			return ErrNotFound
		}
		player.TeamID = *req.TeamID
	}
	return p.playerRepo.Update(player)
}

func (p *playerService) HardDelete(id uint) error {
	if err := p.playerRepo.HardDelete(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFound
		}
		return err
	}
	return nil
}

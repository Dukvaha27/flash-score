package service

import (
	"github.com/Dukvaha27/flash-score/match-service/internal/models"
	"github.com/Dukvaha27/flash-score/match-service/internal/repo"
)

type TeamService interface {
	GetById(id uint) (*models.Team, error)
	GetBySport(id uint) ([]models.Team, error)
	GetList() ([]models.Team, error)
	Create(team models.TeamCreate) (*models.Team, error)
	Update(id uint, team models.TeamUpdate) error
	HardDelete(id uint) error
}

type teamService struct {
	teamRepo  repo.TeamRepo
	sportRepo repo.SportRepo
}

func NewTeamService(team repo.TeamRepo, sport repo.SportRepo) TeamService {
	return &teamService{teamRepo: team, sportRepo: sport}
}

func (t *teamService) GetById(id uint) (*models.Team, error) {
	return t.teamRepo.GetById(id)
}

func (t *teamService) GetBySport(id uint) ([]models.Team, error) {
	if _, err := t.sportRepo.GetById(id); err != nil {
		return nil, ErrNotFound
	}
	return t.teamRepo.GetBySport(id)
}

func (t *teamService) GetList() ([]models.Team, error) {
	return t.teamRepo.GetList()
}

func (t *teamService) Create(team models.TeamCreate) (*models.Team, error) {
	newTeam := models.Team{
		Name:      team.Name,
		City:      team.City,
		SportID:   team.SportID,
		ShortName: team.ShortName,
	}

	if _, err := t.sportRepo.GetById(team.SportID); err != nil {
		return nil, ErrNotFound
	}
	return t.teamRepo.Create(newTeam)
}

func (t *teamService) Update(id uint, team models.TeamUpdate) error {

	existingTeam, err := t.GetById(id)

	if err != nil {
		return ErrNotFound
	}

	if team.Name != nil {
		existingTeam.Name = *team.Name
	}

	if team.ShortName != nil {
		existingTeam.ShortName = *team.ShortName
	}

	if team.City != nil {
		existingTeam.City = *team.City
	}

	if team.SportID != nil {
		if _, err := t.sportRepo.GetById(*team.SportID); err != nil {
			return ErrNotFound
		}

		existingTeam.SportID = *team.SportID
	}

	return t.teamRepo.Update(existingTeam)
}

func (t *teamService) HardDelete(id uint) error {

	if _, err := t.GetById(id); err != nil {
		return ErrNotFound
	}

	return t.teamRepo.HardDelete(id)
}

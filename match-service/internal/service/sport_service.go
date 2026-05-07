package service

import (
	"errors"

	"github.com/Dukvaha27/flash-score/match-service/internal/models"
	"github.com/Dukvaha27/flash-score/match-service/internal/repo"
)

var ErrNotFound = errors.New("record is not found")

type SportService interface {
	Create(name string) (*models.Sport, error)
	Delete(id uint) error
	Update(id uint, name string) error
	GetList() ([]models.Sport, error)
	GetById(id uint) (*models.Sport, error)
}

type sportService struct {
	sportRepo repo.SportRepo
}

func NewSportService(repo repo.SportRepo) SportService {
	return &sportService{sportRepo: repo}
}

func (s *sportService) Create(name string) (*models.Sport, error) {
	sport := models.Sport{Name: name}
	return s.sportRepo.Create(&sport)
}

func (s *sportService) Delete(id uint) error {
	if _, err := s.GetById(id); err != nil {
		return ErrNotFound
	}
	return s.sportRepo.Delete(id)
}

func (s *sportService) Update(id uint, name string) error {
	sport, err := s.sportRepo.GetById(id)

	if err != nil {
		return ErrNotFound
	}

	sport.Name = name

	return s.sportRepo.Update(sport)
}

func (s *sportService) GetList() ([]models.Sport, error) {
	return s.sportRepo.GetList()
}

func (s *sportService) GetById(id uint) (*models.Sport, error) {
	return s.sportRepo.GetById(id)
}

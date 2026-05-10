package service

import (
	"fmt"

	"github.com/Dukvaha27/flash-score/user-service/internal/auth"
	"github.com/Dukvaha27/flash-score/user-service/internal/models"

	"github.com/Dukvaha27/flash-score/user-service/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(req models.LoginRequest) (string, error)
	Register(req models.RegisterRequest) (*models.User, error)
	Update(userID uint, req models.UserUpdate) error
	GetByID(userID uint) (*models.User, error)
	Delete(userID uint) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetByID(userID uint) (*models.User, error) {
	return s.userRepo.GetByID(userID)
}

func (s *userService) Delete(userID uint) error {
	return s.userRepo.Delete(userID)
}

func (s *userService) Update(userID uint, req models.UserUpdate) error {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return err
	}
	if req.Email != nil {

		user.Email = *req.Email
	}

	if req.FavoriteSport != nil {
		user.FavoriteSport = *req.FavoriteSport
	}

	if req.FullName != nil {
		user.FullName = *req.FullName
	}

	if err := s.userRepo.Update(*user); err != nil {
		return err
	}

	return nil

}

func (s *userService) Register(req models.RegisterRequest) (*models.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Ошибка хеширования пароля: %w", err)
	}

	user := models.User{
		FullName:     req.FullName,
		Email:        req.Email,
		PasswordHash: string(hash),
	}

	if err := s.userRepo.Create(&user); err != nil {
		return nil, fmt.Errorf("Ошибка создания пользователя: %w", err)
	}
	return &user, nil
}

func (s *userService) Login(req models.LoginRequest) (string, error) {
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return "", fmt.Errorf("Неверный email или пароль")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return "", fmt.Errorf("Неверный email или пароль")
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return "", fmt.Errorf("Ошибка генерации токена: %w", err)
	}

	return token, nil
}

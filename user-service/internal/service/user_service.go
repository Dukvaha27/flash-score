package service

import (
	"errors"
	"fmt"

	"github.com/Dukvaha27/flash-score/user-service/internal/auth"
	"github.com/Dukvaha27/flash-score/user-service/internal/models"
	"gorm.io/gorm"

	"github.com/Dukvaha27/flash-score/user-service/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var ErrUserNotFound = errors.New("user not found")

type UserService interface {
	Login(req models.LoginRequest) (string, error)
	Register(req models.RegisterRequest) (*models.User, error)
	Update(userID uint, req models.UserUpdate) error
	GetByID(userID uint) (*models.User, error)
	Delete(userID uint) error
}

type userService struct {
	userRepo repository.UserRepository
	secret   string
}

func NewUserService(userRepo repository.UserRepository, secret string) UserService {
	return &userService{userRepo: userRepo, secret: secret}
}

func (s *userService) GetByID(userID uint) (*models.User, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return user, nil
}

func (s *userService) Delete(userID uint) error {
	err := s.userRepo.Delete(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserNotFound
		}
		return err
	}
	return nil
}

func (s *userService) Update(userID uint, req models.UserUpdate) error {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserNotFound
		}
		return err
	}

	if req.Role != nil {
		user.Role = *req.Role
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

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserNotFound
		}
		return err

	}

	return nil

}

func (s *userService) Register(req models.RegisterRequest) (*models.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("ошибка хеширования пароля: %w", err)
	}

	user := models.User{
		FullName:     req.FullName,
		Email:        req.Email,
		PasswordHash: string(hash),
		Role: req.Role,
	}

	if err := s.userRepo.Create(&user); err != nil {
		return nil, fmt.Errorf("ошибка создания пользователя: %w", err)
	}
	return &user, nil
}

func (s *userService) Login(req models.LoginRequest) (string, error) {
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("неверный email или пароль")
		}
		return "", fmt.Errorf("ошибка при поиске пользователя: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return "", errors.New("неверный email или пароль")
	}

	token, err := auth.GenerateToken(user.ID, s.secret)
	if err != nil {
		return "", fmt.Errorf("ошибка генерации токена: %w", err)
	}

	return token, nil
}

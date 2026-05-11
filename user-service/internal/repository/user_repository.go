package repository

import (
	"errors"

	"github.com/Dukvaha27/flash-score/user-service/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByEmail(email string) (*models.User, error)
	GetByID(userID uint) (*models.User, error)
	Create(user *models.User) error
	Update(user models.User) error
	Delete(userID uint) error
}

type gormUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &gormUserRepository{db: db}
}

func (r *gormUserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Такого пользователя не существует")
		} else {
			return nil, err
		}
	}
	return &user, nil
}

func (r *gormUserRepository) GetByID(userID uint) (*models.User, error) {
	var user models.User
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		} else {
			return nil, err
		}
	}
	return &user, nil
}

func (r *gormUserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *gormUserRepository) Update(user models.User) error {
	return r.db.Model(models.User{}).Where("id = ?", user.ID).Updates(user).Error
}

func (r *gormUserRepository) Delete(userID uint) error {
	return r.db.Where("id = ?", userID).Delete(&models.User{}).Error
}

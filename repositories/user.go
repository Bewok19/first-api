package repositories

import (
	"github.com/Bewok19/first-api/entities"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

type UserRepository interface {
	FindAll() ([]entities.User, error)
}

func (r userRepository) FindAll() ([]entities.User, error) {
	var users []entities.User
	err := r.db.Find(&users).Error
	return users, err
}

package user

import (
	"github.com/Bewok19/first-api/entities"
	"github.com/Bewok19/first-api/repositories"
)

type useCase struct {
	userRepo repositories.UserRepository
}

func NewUseCase(userRepo repositories.UserRepository) useCase {
	return useCase{userRepo: userRepo}
}

type UserService interface {
	FindAll() ([]entities.User, error)
}

func (u useCase) FindAll() ([]entities.User, error) {
	return u.userRepo.FindAll()
}

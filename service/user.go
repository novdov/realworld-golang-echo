package service

import (
	"github.com/novdov/realworld-golang-echo/domain"
)

type userService struct {
	repo domain.UserRepository
}

func NewUserService(ur domain.UserRepository) domain.UserService {
	return &userService{
		repo: ur,
	}
}

func (u *userService) GetByEmail(email string) (*domain.User, error) {
	return u.repo.GetByEmail(email)
}

func (u *userService) GetByUsername(username string) (*domain.User, error) {
	return u.repo.GetByUsername(username)
}

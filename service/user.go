package service

import (
	"github.com/novdov/realworld-golang-echo/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userService struct {
	repo domain.UserRepository
}

func NewUserService(ur domain.UserRepository) domain.UserService {
	return &userService{
		repo: ur,
	}
}

func (u *userService) Save(user *domain.User) error {
	return u.repo.Save(user)
}

func (u *userService) Update(user *domain.User) error {
	return u.repo.Update(user)
}

func (u *userService) GetByID(id primitive.ObjectID) (*domain.User, error) {
	return u.repo.GetByID(id)
}

func (u *userService) GetByEmail(email string) (*domain.User, error) {
	return u.repo.GetByEmail(email)
}

func (u *userService) GetByUsername(username string) (*domain.User, error) {
	return u.repo.GetByUsername(username)
}

func (u *userService) FollowUser(user *domain.User, followerID primitive.ObjectID) error {
	return u.repo.FollowUser(user, followerID)
}

func (u *userService) UnFollowUser(user *domain.User, followerID primitive.ObjectID) error {
	return u.repo.UnFollowUser(user, followerID)
}

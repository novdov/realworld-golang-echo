package domain

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string
	Email    string
	Password string
	Bio      string
	Image    string
	Follows  []primitive.ObjectID
}

func (u *User) Following() bool {
	return len(u.Follows) > 0
}

func (u *User) HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}

func (u *User) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
	return err == nil
}

type UserRepository interface {
	Save(*User) error
	GetByID(primitive.ObjectID) (*User, error)
	GetByEmail(string) (*User, error)
	GetByUsername(string) (*User, error)
	Update(*User) error
	FollowUser(*User, primitive.ObjectID) error
	UnFollowUser(*User, primitive.ObjectID) error
}

type UserService interface {
	Save(*User) error
	GetByID(primitive.ObjectID) (*User, error)
	GetByEmail(string) (*User, error)
	GetByUsername(string) (*User, error)
	Update(*User) error
	FollowUser(*User, primitive.ObjectID) error
	UnFollowUser(*User, primitive.ObjectID) error
}

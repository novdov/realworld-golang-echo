package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID         primitive.ObjectID
	Username   string
	Email      string
	Password   string
	Bio        string
	Image      string
	Followers  []User
	Followings []User
}

func (u *User) Following() bool {
	return len(u.Followings) > 0
}

type UserRepository interface {
	GetByEmail(email string) (*User, error)
	GetByUsername(username string) (*User, error)
}

type UserService interface {
	GetByEmail(email string) (*User, error)
	GetByUsername(username string) (*User, error)
}

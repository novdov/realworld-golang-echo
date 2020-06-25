package user

import (
	"context"

	"github.com/novdov/realworld-golang-echo/domain"
	"github.com/novdov/realworld-golang-echo/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db             *mongo.Database
	collectionName string
}

func (u *userRepository) collection() *mongo.Collection {
	return u.db.Collection(u.collectionName)
}

func NewUserRepository(db *mongo.Database, collectionName string) domain.UserRepository {
	return &userRepository{
		db:             db,
		collectionName: collectionName,
	}
}

func (u *userRepository) Save(user *domain.User) error {
	if user.ID == primitive.NilObjectID {
		user.ID = primitive.NewObjectID()
	}
	_, err := u.collection().InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) Update(user *domain.User) error {
	doc, err := utils.ToDocument(user)
	if err != nil {
		return err
	}
	_, err = u.collection().UpdateOne(
		context.TODO(),
		bson.D{{"_id", user.ID}},
		bson.D{{"$set", doc}},
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) GetByID(id primitive.ObjectID) (*domain.User, error) {
	return u.getUser("_id", id)
}

func (u *userRepository) GetByEmail(email string) (*domain.User, error) {
	return u.getUser("email", email)
}

func (u *userRepository) GetByUsername(username string) (*domain.User, error) {
	return u.getUser("username", username)
}

func (u *userRepository) FollowUser(user *domain.User, followerID primitive.ObjectID) error {
	var exists bool
	for _, followedID := range user.Follows {
		if followedID == followerID {
			exists = true
		}
	}
	if !exists {
		user.Follows = append(user.Follows, followerID)
	}
	return u.Update(user)
}

func (u *userRepository) UnFollowUser(user *domain.User, followerID primitive.ObjectID) error {
	notFound := -1
	idx := notFound
	for i, followedID := range user.Follows {
		if followedID == followerID {
			idx = i
		}
	}
	if idx > notFound {
		user.Follows = append(user.Follows[:idx], user.Follows[(idx+1):]...)
	}
	return u.Update(user)
}

func (u *userRepository) getUser(key string, value interface{}) (*domain.User, error) {
	result := u.collection().FindOne(
		context.TODO(),
		bson.D{{key, value}},
	)

	var user domain.User
	if err := result.Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

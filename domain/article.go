package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	ID          primitive.ObjectID   `bson:"_id"`
	Slug        string               `bson:"slug"`
	Title       string               `bson:"title"`
	Description string               `bson:"description"`
	Body        string               `bson:"body"`
	TagList     []string             `bson:"tagList"`
	CreatedAt   time.Time            `bson:"createdAt"`
	UpdatedAt   time.Time            `bson:"updatedAt"`
	Author      primitive.ObjectID   `bson:"author"`
	Favorites   []primitive.ObjectID `bson:"favorites"`
}

type Comment struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
	Body      string             `bson:"body"`
	Author    primitive.ObjectID `bson:"author"`
}

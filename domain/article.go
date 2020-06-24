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

type ArticleRepository interface {
	ListByTag(tag string) ([]*Article, error)
	ListByAuthor(username string) ([]*Article, error)
	ListByLimit(limit int) ([]*Article, error)
	ListByOffset(offset int) ([]*Article, error)
	GetBySlug(slug string) (*Article, error)
	Save(article *Article) error
	Update(article *Article) error
	Delete(article *Article) error
}

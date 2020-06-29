package domain

import (
	"strings"
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
	Comments    []*Comment           `bson:"comments"`
}

func (a *Article) UpdateSlug() {
	title := strings.ToLower(a.Title)
	titleSplit := strings.Split(title, " ")
	a.Slug = strings.Join(titleSplit, "-")
}

func (a *Article) AddComments(comment *Comment) {
	a.Comments = append(a.Comments, comment)
}

type Comment struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
	Body      string             `bson:"body"`
	Author    primitive.ObjectID `bson:"author"`
}

type ArticleRepository interface {
	Find(query map[string]string, skip int64, limit int64) ([]*Article, error)
	GetBySlug(slug string) (*Article, error)
	Save(article *Article) error
	Update(article *Article) error
	Delete(article *Article) error
	GetTags() ([]interface{}, error)
	AddComments(article *Article, comment *Comment) error
	DeleteComments(article *Article, id primitive.ObjectID) error
}

type ArticleService interface {
	Save(article *Article) error
	Update(article *Article) error
	Delete(article *Article) error
	GetBySlug(slug string) (*Article, error)
	GetTags() ([]interface{}, error)
	AddComments(article *Article, comment *Comment) error
	DeleteComments(article *Article, id primitive.ObjectID) error
}

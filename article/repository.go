package article

import (
	"context"

	"github.com/novdov/realworld-golang-echo/domain"
	"github.com/novdov/realworld-golang-echo/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type articleRepository struct {
	db             *mongo.Database
	collectionName string
}

func (a *articleRepository) collection() *mongo.Collection {
	return a.db.Collection(a.collectionName)
}

func NewArticleRepository(db *mongo.Database, collectionName string) domain.ArticleRepository {
	return &articleRepository{
		db:             db,
		collectionName: collectionName,
	}
}

func (a *articleRepository) GetBySlug(slug string) (*domain.Article, error) {
	return a.getArticle("slug", slug)
}

func (a *articleRepository) Save(article *domain.Article) error {
	if article.ID == primitive.NilObjectID {
		article.ID = primitive.NewObjectID()
	}

	_, err := a.collection().InsertOne(context.TODO(), article)
	if err != nil {
		return err
	}
	return nil
}

func (a *articleRepository) Update(article *domain.Article) error {
	doc, err := utils.ToDocument(article)
	if err != nil {
		return err
	}
	_, err = a.collection().UpdateOne(
		context.TODO(),
		bson.D{{"_id", article.ID}},
		bson.D{{"$set", doc}},
	)
	if err != nil {
		return err
	}
	return nil
}

func (a *articleRepository) Delete(article *domain.Article) error {
	_, err := a.collection().DeleteOne(
		context.TODO(),
		bson.D{{"_id", article.ID}},
	)
	if err != nil {
		return err
	}
	return nil
}

func (a *articleRepository) getArticle(key string, value interface{}) (*domain.Article, error) {
	result := a.collection().FindOne(
		context.TODO(),
		bson.D{{key, value}},
	)

	var article domain.Article
	if err := result.Decode(&article); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &article, nil
}

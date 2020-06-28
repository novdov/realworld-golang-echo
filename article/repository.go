package article

import (
	"context"
	"time"

	"github.com/novdov/realworld-golang-echo/domain"
	"github.com/novdov/realworld-golang-echo/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (a *articleRepository) Find(query map[string]string, skip int64, limit int64) ([]*domain.Article, error) {
	var result []*domain.Article

	var bsonFilter = bson.D{}
	for key, value := range query {
		bsonFilter = append(bsonFilter, bson.E{Key: key, Value: value})
	}

	cur, err := a.collection().Find(
		context.TODO(),
		bsonFilter,
		getFindOptions(skip, limit),
	)

	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var article *domain.Article
		err := cur.Decode(&article)
		if err != nil {
			return nil, err
		}
		result = append(result, article)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(context.TODO())
	return result, nil
}

func (a *articleRepository) GetBySlug(slug string) (*domain.Article, error) {
	return a.getArticle("slug", slug)
}

func (a *articleRepository) Save(article *domain.Article) error {
	if article.ID == primitive.NilObjectID {
		article.ID = primitive.NewObjectID()
	}
	if article.CreatedAt.IsZero() {
		article.CreatedAt = currentTime()
	}
	article.UpdatedAt = currentTime()
	article.UpdateSlug()

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

func (a *articleRepository) GetTags() ([]interface{}, error) {
	result, err := a.collection().Distinct(
		context.TODO(),
		"tagList",
		bson.D{{}},
		options.Distinct().SetMaxTime(10*time.Second),
	)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *articleRepository) AddComments(article *domain.Article, comment *domain.Comment) error {
	if comment.ID == primitive.NilObjectID {
		comment.ID = primitive.NewObjectID()
	}
	article.AddComments(comment)
	return a.Update(article)
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

func getFindOptions(skip int64, limit int64) *options.FindOptions {
	opts := &options.FindOptions{}
	if skip != 0 {
		opts.Skip = &skip
	}
	if limit != 0 {
		opts.Limit = &limit
	}
	return opts
}

func currentTime() time.Time {
	return time.Now().Local()
}

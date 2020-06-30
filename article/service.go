package article

import (
	"github.com/novdov/realworld-golang-echo/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type articleService struct {
	repo domain.ArticleRepository
}

func NewArticleService(ar domain.ArticleRepository) domain.ArticleService {
	return &articleService{
		repo: ar,
	}
}

func (a *articleService) Find(query map[string]string, skip int64, limit int64) ([]*domain.Article, error) {
	return a.repo.Find(query, skip, limit)
}

func (a *articleService) Save(article *domain.Article) error {
	return a.repo.Save(article)
}

func (a *articleService) GetBySlug(slug string) (*domain.Article, error) {
	return a.repo.GetBySlug(slug)
}

func (a *articleService) Update(article *domain.Article) error {
	return a.repo.Update(article)
}

func (a *articleService) Delete(article *domain.Article) error {
	return a.repo.Delete(article)
}

func (a *articleService) GetTags() ([]interface{}, error) {
	return a.repo.GetTags()
}

func (a *articleService) AddComments(article *domain.Article, comment *domain.Comment) error {
	return a.repo.AddComments(article, comment)
}

func (a *articleService) DeleteComments(article *domain.Article, id primitive.ObjectID) error {
	return a.repo.DeleteComments(article, id)
}

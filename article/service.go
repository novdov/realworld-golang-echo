package article

import "github.com/novdov/realworld-golang-echo/domain"

type articleService struct {
	repo domain.ArticleRepository
}

func NewArticleService(ar domain.ArticleRepository) domain.ArticleService {
	return &articleService{
		repo: ar,
	}
}

func (a *articleService) Save(article *domain.Article) error {
	return a.repo.Save(article)
}

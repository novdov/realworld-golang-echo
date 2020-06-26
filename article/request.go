package article

import (
	"github.com/labstack/echo/v4"
	"github.com/novdov/realworld-golang-echo/domain"
)

type articleCreateRequest struct {
	Article struct {
		Title       string   `json:"title" validate:"required"`
		Description string   `json:"description" validate:"description"`
		Body        string   `json:"body" validate:"required"`
		TagList     []string `json:"tagList"`
	} `json:"article"`
}

func (a *articleCreateRequest) bind(c echo.Context, article *domain.Article) error {
	if err := c.Bind(a); err != nil {
		return err
	}
	if err := c.Validate(a); err != nil {
		return err
	}

	article.Title = a.Article.Title
	article.Description = a.Article.Description
	article.Body = a.Article.Body
	article.TagList = a.Article.TagList
	return nil
}

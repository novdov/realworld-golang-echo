package article

import (
	"github.com/labstack/echo/v4"
	"github.com/novdov/realworld-golang-echo/domain"
)

type articleCreateRequest struct {
	Article struct {
		Title       string   `json:"title" validate:"required"`
		Description string   `json:"description" validate:"required"`
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

type articleUpdateRequest struct {
	Article struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Body        string `json:"body"`
	} `json:"article"`
}

func (a *articleUpdateRequest) bind(c echo.Context, article *domain.Article) error {
	if err := c.Bind(a); err != nil {
		return err
	}
	if err := c.Validate(a); err != nil {
		return err
	}

	if a.Article.Title != article.Title {
		article.Title = a.Article.Title
		article.UpdateSlug()
	}

	if a.Article.Description != article.Description {
		article.Description = a.Article.Description
	}

	if a.Article.Body != article.Body {
		article.Body = a.Article.Body
	}

	article.UpdatedAt = currentTime()
	return nil
}

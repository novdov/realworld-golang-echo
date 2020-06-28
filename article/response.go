package article

import (
	"time"

	"github.com/novdov/realworld-golang-echo/domain"
)

type singleArticleResponse struct {
	Article struct {
		Title          string    `json:"title"`
		Slug           string    `json:"slug"`
		Description    string    `json:"description"`
		Body           string    `json:"body"`
		TagList        []string  `json:"tagList"`
		CreatedAt      time.Time `json:"createdAt"`
		UpdatedAt      time.Time `json:"updatedAt"`
		Favorited      bool      `json:"favorited"`
		FavoritesCount int64     `json:"favoritesCount"`
		Author         struct {
			Username  string `json:"username"`
			Bio       string `json:"bio"`
			Image     string `json:"image"`
			Following bool   `json:"following"`
		} `json:"author"`
	} `json:"article"`
}

func newSingleArticleResponse(article *domain.Article, user *domain.User) *singleArticleResponse {
	resp := &singleArticleResponse{}
	resp.Article.Title = article.Title
	resp.Article.Slug = article.Slug
	resp.Article.Description = article.Description
	resp.Article.Body = article.Body
	resp.Article.TagList = article.TagList
	resp.Article.CreatedAt = article.CreatedAt
	resp.Article.UpdatedAt = article.UpdatedAt
	resp.Article.Author.Username = user.Username
	resp.Article.Author.Bio = user.Bio
	resp.Article.Author.Image = user.Image
	resp.Article.Author.Following = user.Following()
	return resp
}

type singleCommentResponse struct {
	Comment struct {
		Body string `json:"body"`
	} `json:"comment"`
}

func newSingleCommentResponse(comment *domain.Comment) *singleCommentResponse {
	resp := &singleCommentResponse{}
	resp.Comment.Body = comment.Body
	return resp
}

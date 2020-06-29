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
		ID        string    `json:"id"`
		Body      string    `json:"body"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		Author    struct {
			Username  string `json:"username"`
			Bio       string `json:"bio"`
			Image     string `json:"image"`
			Following bool   `json:"following"`
		} `json:"author"`
	} `json:"comment"`
}

func newSingleCommentResponse(comment *domain.Comment, user *domain.User) *singleCommentResponse {
	resp := &singleCommentResponse{}
	resp.Comment.ID = comment.ID.Hex()
	resp.Comment.Body = comment.Body
	resp.Comment.CreatedAt = comment.CreatedAt
	resp.Comment.UpdatedAt = comment.UpdatedAt
	resp.Comment.Author.Username = user.Username
	resp.Comment.Author.Bio = user.Bio
	resp.Comment.Author.Image = user.Image
	resp.Comment.Author.Following = user.Following()
	return resp
}

type multipleCommentsResponse struct {
	Comments []*singleCommentResponse `json:"comments"`
}

func newMultipleCommentsResponse(comments []*singleCommentResponse) *multipleCommentsResponse {
	return &multipleCommentsResponse{Comments: comments}
}

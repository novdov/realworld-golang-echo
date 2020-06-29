package article

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/novdov/realworld-golang-echo/domain"
	"github.com/novdov/realworld-golang-echo/errors"
	"github.com/novdov/realworld-golang-echo/utils"
)

type Handler struct {
	articleService domain.ArticleService
	userService    domain.UserService
}

func NewHandler(as domain.ArticleService, us domain.UserService) *Handler {
	return &Handler{articleService: as, userService: us}
}

func (h *Handler) Register(g *echo.Group) {
	jwtMiddleware := middleware.JWT(utils.JWTSecret)

	article := g.Group("/articles", jwtMiddleware)
	article.POST("", h.Create)
	article.GET("/:slug", h.GetSingleArticle)
	article.PUT("/:slug", h.Update)
	article.DELETE("/:slug", h.Delete)
	article.POST("/:slug/comments", h.AddComments)

	tags := g.Group("/tags")
	tags.GET("", h.GetTags)
}

func (h *Handler) Create(c echo.Context) error {
	user, err := h.userService.GetByID(utils.GetUserIDFromJWT(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.NewError(err))
	}
	if user == nil {
		return c.JSON(http.StatusNotFound, errors.NewError(errors.NotFound))
	}

	article := &domain.Article{
		Author: user.ID,
	}
	req := &articleCreateRequest{}

	if err := req.bind(c, article); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errors.NewError(err))
	}
	if err := h.articleService.Save(article); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errors.NewError(err))
	}
	return c.JSON(http.StatusCreated, newSingleArticleResponse(article, user))
}

func (h *Handler) GetSingleArticle(c echo.Context) error {
	slug := c.Param("slug")
	article, err := h.articleService.GetBySlug(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.NewError(err))
	}
	if article == nil {
		return c.JSON(http.StatusNotFound, errors.NewError(errors.NotFound))
	}

	user, err := h.userService.GetByID(article.Author)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.NewError(err))
	}
	if user == nil {
		return c.JSON(http.StatusNotFound, errors.NewError(errors.NotFound))
	}

	return c.JSON(http.StatusOK, newSingleArticleResponse(article, user))
}

func (h *Handler) Update(c echo.Context) error {
	slug := c.Param("slug")
	article, err := h.articleService.GetBySlug(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.NewError(err))
	}
	if article == nil {
		return c.JSON(http.StatusNotFound, errors.NewError(errors.NotFound))
	}

	req := &articleUpdateRequest{}

	if err := req.bind(c, article); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errors.NewError(err))
	}
	if err := h.articleService.Update(article); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errors.NewError(err))
	}

	user, err := h.userService.GetByID(article.Author)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.NewError(err))
	}
	if user == nil {
		return c.JSON(http.StatusNotFound, errors.NewError(errors.NotFound))
	}

	return c.JSON(http.StatusOK, newSingleArticleResponse(article, user))
}

func (h *Handler) Delete(c echo.Context) error {
	slug := c.Param("slug")
	article, err := h.articleService.GetBySlug(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.NewError(err))
	}
	if article == nil {
		return c.JSON(http.StatusNotFound, errors.NewError(errors.NotFound))
	}

	req := &articleUpdateRequest{}

	if err := req.bind(c, article); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errors.NewError(err))
	}
	if err := h.articleService.Delete(article); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errors.NewError(err))
	}

	user, err := h.userService.GetByID(article.Author)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.NewError(err))
	}
	if user == nil {
		return c.JSON(http.StatusNotFound, errors.NewError(errors.NotFound))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"result": "deleted article"})
}

func (h *Handler) GetTags(c echo.Context) error {
	tags, err := h.articleService.GetTags()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.NewError(err))
	}
	return c.JSON(http.StatusOK, map[string][]interface{}{"tags": tags})
}

func (h *Handler) AddComments(c echo.Context) error {
	slug := c.Param("slug")
	article, err := h.articleService.GetBySlug(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.NewError(err))
	}
	if article == nil {
		return c.JSON(http.StatusNotFound, errors.NewError(errors.NotFound))
	}

	user, err := h.userService.GetByID(article.Author)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.NewError(err))
	}
	if user == nil {
		return c.JSON(http.StatusNotFound, errors.NewError(errors.NotFound))
	}

	req := &commentsCreateRequest{}
	var comment domain.Comment
	if err := req.bind(c, &comment, user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errors.NewError(err))
	}

	if err := h.articleService.AddComments(article, &comment); err != nil {
		return c.JSON(http.StatusInternalServerError, errors.NewError(err))
	}
	if err := h.articleService.Update(article); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errors.NewError(err))
	}

	return c.JSON(http.StatusCreated, newSingleCommentResponse(&comment, user))
}

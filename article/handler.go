package article

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/novdov/realworld-golang-echo/domain"
	"github.com/novdov/realworld-golang-echo/errors"
)

type Handler struct {
	articleService domain.ArticleService
}

func NewHandler(as domain.ArticleService) *Handler {
	return &Handler{articleService: as}
}

func (h *Handler) Create(c echo.Context) error {
	var article domain.Article
	req := &articleCreateRequest{}

	if err := req.bind(c, &article); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errors.NewError(err))
	}
	if err := h.articleService.Save(&article); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errors.NewError(err))
	}
	return c.JSON(http.StatusCreated, new)
}

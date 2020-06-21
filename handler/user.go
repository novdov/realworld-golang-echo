package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/novdov/realworld-golang-echo/domain"
)

type UserHandler struct {
	userService domain.UserService
}

func NewUserHandler(g *echo.Group, us domain.UserService) {
	h := &UserHandler{
		userService: us,
	}
	profiles := g.Group("/profiles")
	profiles.GET("/:username", h.GetProfile)

	users := g.Group("/users")
	users.POST("", h.Signup)
}

func (h *UserHandler) Signup(c echo.Context) error {
	var u domain.User
	req := userRegisterRequest{}

	if err := req.bind(c, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	if err := h.userService.Save(&u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	return c.JSON(http.StatusCreated, newUserResponse(&u))
}

func (h *UserHandler) GetProfile(c echo.Context) error {
	username := c.Param("username")
	u, err := h.userService.GetByUsername(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, NotFound())
	}
	return c.JSON(http.StatusOK, newProfileResponse(u))
}

package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/novdov/realworld-golang-echo/domain"
	"github.com/novdov/realworld-golang-echo/utils"
)

type UserHandler struct {
	userService domain.UserService
}

func NewUserHandler(g *echo.Group, us domain.UserService) {
	h := &UserHandler{
		userService: us,
	}

	jwtMiddleware := middleware.JWT(utils.JWTSecret)

	profiles := g.Group("/profiles", jwtMiddleware)
	profiles.GET("/:username", h.GetProfile)

	guestUsers := g.Group("/users")
	guestUsers.POST("", h.Signup)
	guestUsers.POST("/login", h.Login)
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

func (h *UserHandler) Login(c echo.Context) error {
	req := userLoginRequest{}
	if err := req.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}

	u, err := h.userService.GetByEmail(req.User.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusForbidden, AccessForbidden())
	}

	if !u.CheckPassword(req.User.Password) {
		return c.JSON(http.StatusForbidden, AccessForbidden())
	}

	return c.JSON(http.StatusOK, newUserResponse(u))
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

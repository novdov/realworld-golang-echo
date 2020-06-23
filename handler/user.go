package handler

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/novdov/realworld-golang-echo/domain"
	"github.com/novdov/realworld-golang-echo/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	userService domain.UserService
}

func NewUserHandler(us domain.UserService) *UserHandler {
	return &UserHandler{userService: us}
}

func (h *UserHandler) Register(g *echo.Group) {
	jwtMiddleware := middleware.JWT(utils.JWTSecret)

	profile := g.Group("/profiles", jwtMiddleware)
	profile.GET("/:username", h.GetProfile)
	profile.POST("/:username/follow", h.Follow)
	profile.DELETE("/:username/follow", h.UnFollow)

	auth := g.Group("/users")
	auth.POST("", h.Signup)
	auth.POST("/login", h.Login)

	user := g.Group("/user", jwtMiddleware)
	user.GET("", h.GetCurrentUser)
	user.PUT("", h.UpdateUser)
}

func (h *UserHandler) Signup(c echo.Context) error {
	var u domain.User
	req := &userRegisterRequest{}

	if err := req.bind(c, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	if err := h.userService.Save(&u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	return c.JSON(http.StatusCreated, newUserResponse(&u))
}

func (h *UserHandler) Login(c echo.Context) error {
	req := &userLoginRequest{}
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

func (h *UserHandler) GetCurrentUser(c echo.Context) error {
	id := getIDFromToken(c)
	u, err := h.userService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, NotFound())
	}
	return c.JSON(http.StatusOK, newUserResponse(u))
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	id := getIDFromToken(c)
	u, err := h.userService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, NotFound())
	}

	req := &userUpdateRequest{}
	if err := req.bind(c, u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	if err := h.userService.Update(u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	return c.JSON(http.StatusOK, newUserResponse(u))
}

func (h *UserHandler) Follow(c echo.Context) error {
	id := getIDFromToken(c)
	u, err := h.userService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, NotFound())
	}

	username := c.Param("username")
	follower, err := h.userService.GetByUsername(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	if follower == nil {
		return c.JSON(http.StatusNotFound, NotFound())
	}

	err = h.userService.FollowUser(u, follower.ID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	return c.JSON(http.StatusOK, newProfileResponse(u))
}

func (h *UserHandler) UnFollow(c echo.Context) error {
	id := getIDFromToken(c)
	u, err := h.userService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, NotFound())
	}

	username := c.Param("username")
	follower, err := h.userService.GetByUsername(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	if follower == nil {
		return c.JSON(http.StatusNotFound, NotFound())
	}

	err = h.userService.UnFollowUser(u, follower.ID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	return c.JSON(http.StatusOK, newProfileResponse(u))
}

func getIDFromToken(c echo.Context) primitive.ObjectID {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	idStr, ok := claims["id"].(string)
	if !ok {
		return primitive.NilObjectID
	}
	id, _ := primitive.ObjectIDFromHex(idStr)
	return id
}

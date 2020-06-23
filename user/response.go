package user

import (
	"github.com/labstack/echo/v4"
	"github.com/novdov/realworld-golang-echo/domain"
	"github.com/novdov/realworld-golang-echo/utils"
)

type profileResponse struct {
	Profile struct {
		Username  string `json:"username"`
		Bio       string `json:"bio"`
		Image     string `json:"image"`
		Following bool   `json:"following"`
	} `json:"profile"`
}

type userResponse struct {
	User struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Bio      string `json:"bio"`
		Image    string `json:"image"`
		Token    string `json:"token"`
	} `json:"user"`
}

func newProfileResponse(user *domain.User) *profileResponse {
	resp := &profileResponse{}
	resp.Profile.Username = user.Username
	resp.Profile.Bio = user.Bio
	resp.Profile.Image = user.Image
	resp.Profile.Following = user.Following()
	return resp
}

func newUserResponse(user *domain.User) *userResponse {
	resp := &userResponse{}
	resp.User.Username = user.Username
	resp.User.Email = user.Email
	resp.User.Bio = user.Bio
	resp.User.Image = user.Image
	resp.User.Token, _ = utils.GenerateJWT(user.Email, user.ID)
	return resp
}

type ResponseError struct {
	Errors map[string]interface{}
}

func NewError(err error) ResponseError {
	e := ResponseError{}
	e.Errors = make(map[string]interface{})
	switch v := err.(type) {
	case *echo.HTTPError:
		e.Errors["body"] = v.Message
	default:
		e.Errors["body"] = v.Error()
	}
	return e
}

func NotFound() ResponseError {
	e := ResponseError{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = "resource not found"
	return e
}

func AccessForbidden() ResponseError {
	e := ResponseError{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = "access forbidden"
	return e
}

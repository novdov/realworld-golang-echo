package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/novdov/realworld-golang-echo/domain"
)

type profileResponse struct {
	Profile struct {
		Username  string `json:"username"`
		Bio       string `json:"bio"`
		Image     string `json:"image"`
		Following bool   `json:"following"`
	} `json:"profile"`
}

func newProfileResponse(user *domain.User) *profileResponse {
	resp := &profileResponse{}
	resp.Profile.Username = user.Username
	resp.Profile.Bio = user.Bio
	resp.Profile.Image = user.Image
	resp.Profile.Following = user.Following()
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

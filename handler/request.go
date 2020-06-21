package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/novdov/realworld-golang-echo/domain"
)

type userRegisterRequest struct {
	User struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

func (u *userRegisterRequest) bind(c echo.Context, user *domain.User) error {
	if err := c.Bind(&u); err != nil {
		return err
	}
	if err := c.Validate(u); err != nil {
		return err
	}

	user.Username = u.User.Username
	user.Email = u.User.Email

	h, err := user.HashPassword(u.User.Password)
	if err != nil {
		return err
	}
	user.Password = h
	return nil
}

type userLoginRequest struct {
	User struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

func (u *userLoginRequest) bind(c echo.Context) error {
	if err := c.Bind(&u); err != nil {
		return err
	}
	if err := c.Validate(u); err != nil {
		return err
	}
	return nil
}

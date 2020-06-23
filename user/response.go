package user

import (
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

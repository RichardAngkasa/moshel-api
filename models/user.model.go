package models

type UserDataInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

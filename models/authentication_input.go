package models

type AuthenticationInput struct {
	Email    string `json: "email" binding: "required"`
	Password string `json: "password" binding: "required"`
	UserType string `json: "userType" binding: "required`
}

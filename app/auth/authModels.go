package app_auth

import (
	app_users "github.com/DenysNykoriak/go-api/app/users"
	"github.com/DenysNykoriak/go-api/core"
	"github.com/go-playground/validator"
)

type SignUpBody struct {
	FirstName string `json:"first_name,omitempty" validate:"required"`
	LastName  string `json:"last_name,omitempty" validate:"required"`
	Email     string `json:"email,omitempty" validate:"required,email"`
	Password  string `json:"password,omitempty" validate:"required"`
}

func (body *SignUpBody) Validate() (bool, SignUpBody) {

	err := core.Validate.Struct(body)

	//Email candidate
	userCandidate := core.PostgresDB.Model(&app_users.User{}).Where("email = ?", body.Email).First(&app_users.User{})
	if userCandidate.RowsAffected > 0 {
		return false, SignUpBody{Email: "taken"}
	}

	if err == nil {
		return true, *body
	}

	errs := err.(validator.ValidationErrors)

	errorResponse := SignUpBody{}

	for _, e := range errs {
		switch e.Field() {
		case "FirstName":
			errorResponse.FirstName = e.Tag()
		case "LastName":
			errorResponse.LastName = e.Tag()
		case "Email":
			errorResponse.Email = e.Tag()
		case "Password":
			errorResponse.Password = e.Tag()
		}
	}

	return false, errorResponse
}

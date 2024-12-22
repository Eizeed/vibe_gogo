package forms

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type UserForm struct {};

type RegisterForm struct {
    Email           string      `json:"email" binding:"required,email"`
    Username        string      `json:"username" binding:"required,min=3"`
    Fullname        string      `json:"fullname" binding:"required,min=2"`
    Password        string      `json:"password" binding:"required,min=3"`
}

type LoginForm struct {
    Email           string      `json:"email" binding:"required,email"`
    Password        string      `json:"password" binding:"required,min=3"`
}

type UpdateForm struct {
    Email           string      `json:"email"`
    Username        string      `json:"username"`
    Fullname        string      `json:"fullname"`
    Password        string      `json:"password"`
}

func (f UserForm) Fullname(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter your name"
		}
		return errMsg[0]
	case "min", "max":
		return "Your name should be at least 2 characters long"
	default:
		return "Something went wrong, please try again later"
	}
}

//Email ...
func (f UserForm) Email(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter your email"
		}
		return errMsg[0]
	case "min", "max", "email":
		return "Please enter a valid email"
	default:
		return "Something went wrong, please try again later"
	}
}

//Password ...
func (f UserForm) Password(tag string) (message string) {
	switch tag {
	case "required":
		return "Please enter your password"
	case "min", "max":
		return "Your password should be more than 3 characters"
	default:
		return "Something went wrong, please try again later"
	}
}

//Signin ...
func (f UserForm) Login(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Email" {
				return f.Email(err.Tag())
			}
			if err.Field() == "Password" {
				return f.Password(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}

//Register ...
func (f UserForm) Register(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Name" {
				return f.Fullname(err.Tag())
			}

			if err.Field() == "Email" {
				return f.Email(err.Tag())
			}

			if err.Field() == "Password" {
				return f.Password(err.Tag())
			}

		}
	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}

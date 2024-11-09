package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	gorm.Model
	FullName         string    `json:"fullName" validate:"required"`
	Username         string    `json:"username" validate:"required"`
	Email            string    `json:"email"`
	Bio              string    `json:"bio"`
	ProfilePic       string    `json:"profilePic"`
	Verified         bool      `json:"verified"`
	Role             Role      `json:"role"`
	Followers        string    `json:"followers"`
	Following        string    `json:"following"`
	LoggedWithGoogle bool      `json:"loggedWithGoogle"`
	LastActive       time.Time `json:"lastActive"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (u *User) Validate() error {
	err := validate.Struct(u)

	if err != nil {
		validateErrors, ok := err.(validator.ValidationErrors)

		if ok {
			errorMessage := make(map[string]string)

			for _, err := range validateErrors {

				switch err.Field() {
				case "FullName":
					errorMessage["fullName"] = "Full Name is required"
				case "Username":
					errorMessage["username"] = "Username is required"
				}

			}

			return errorMessage

		}
		return err

	}

	return nil
}

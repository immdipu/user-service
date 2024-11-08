package models

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	gorm.Model
	FullName         string    `json:"fullName"`
	Username         string    `json:"username"`
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

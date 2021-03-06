package models

import (
	"time"
)

type Users struct {
	Id            int    `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Authorization string
}

type DisplayUser struct {
	Id            int       `json:"id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Email         string    `json:"email"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Authorization string    `json:"-"`
}

type GoogleObject struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Verified  bool
	Name      string `json:"name"`
	GivenName string
	Picture   string
	Locale    string
}

type Password struct {
	Password string `json:"password"`
}

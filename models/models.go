package models

import "errors"

// maybe better place for it?
var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
	ErrDuplicateUsername  = errors.New("models: duplicate username")
	ErrTooManySpaces      = errors.New("inupt data: too many spaces in field")
	ErrUserNotActivated   = errors.New("models: user is not activated")
)

type User struct {
	Id             int    `json:"id,omitempty"`
	FirstName      string `json:"firstname,omitempty"`
	LastName       string `json:"lastname,omitempty"`
	Email          string `json:"email,omitempty"`
	Password       string `json:"password,omitempty"`
	HashedPassword []byte `json:"-"`
	// IsAdmin        int    `json:"is_admin,omitempty"`
	// Activated      int    `json:"activated,omitempty"`
}

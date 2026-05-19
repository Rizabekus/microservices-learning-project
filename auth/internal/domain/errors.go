package domain

import "errors"

var (
	ErrInvalidFirstName    = errors.New("invalid first name")
	ErrInvalidLastName     = errors.New("invalid last name")
	ErrInvalidEmail        = errors.New("invalid email")
	ErrInvalidPassword     = errors.New("invalid password")
	ErrInvalidMobileNumber = errors.New("invalid mobile number")
)

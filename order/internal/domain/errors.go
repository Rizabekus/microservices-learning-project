package domain

import "errors"

var (
	ErrInvalidUserID   = errors.New("invalid user id")
	ErrInvalidAmount   = errors.New("invalid amount")
	ErrInvalidCurrency = errors.New("invalid currency")
)

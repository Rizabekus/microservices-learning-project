package domain

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string
	MobileNumber string
	CreatedAt    time.Time
}

func NewUser(firstName string, lastName string, email string, password string, mobileNumber string) (*User, error) {
	if err := validateFirstName(firstName); err != nil {
		return nil, err
	}

	if err := validateLastName(lastName); err != nil {
		return nil, err
	}

	if err := validateEmail(email); err != nil {
		return nil, err
	}

	if err := validatePassword(password); err != nil {
		return nil, err
	}

	if err := validateMobileNumber(mobileNumber); err != nil {
		return nil, err
	}

	return &User{
		ID:           uuid.New(),
		FirstName:    strings.TrimSpace(firstName),
		LastName:     strings.TrimSpace(lastName),
		Email:        strings.ToLower(strings.TrimSpace(email)),
		PasswordHash: password,
		MobileNumber: strings.TrimSpace(mobileNumber),
		CreatedAt:    time.Now().UTC(),
	}, nil
}

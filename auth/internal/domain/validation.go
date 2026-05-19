package domain

import (
	"net/mail"
	"regexp"
	"strings"
	"unicode"
)

var phoneRegex = regexp.MustCompile(`^\+?[1-9]\d{7,14}$`)

func validateFirstName(firstName string) error {
	firstName = strings.TrimSpace(firstName)

	if len(firstName) < 2 || len(firstName) > 50 {
		return ErrInvalidFirstName
	}

	for _, r := range firstName {
		if !unicode.IsLetter(r) && r != ' ' && r != '-' {
			return ErrInvalidFirstName
		}
	}

	return nil
}

func validateLastName(lastName string) error {
	lastName = strings.TrimSpace(lastName)

	if len(lastName) < 2 || len(lastName) > 50 {
		return ErrInvalidLastName
	}

	for _, r := range lastName {
		if !unicode.IsLetter(r) && r != ' ' && r != '-' {
			return ErrInvalidLastName
		}
	}

	return nil
}

func validateEmail(email string) error {
	email = strings.TrimSpace(email)

	if len(email) == 0 || len(email) > 255 {
		return ErrInvalidEmail
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return ErrInvalidEmail
	}

	return nil
}

func validatePassword(password string) error {
	if len(password) < 8 || len(password) > 72 {
		return ErrInvalidPassword
	}

	var hasUpper bool
	var hasLower bool
	var hasDigit bool

	for _, r := range password {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasDigit = true
		}
	}

	if !hasUpper || !hasLower || !hasDigit {
		return ErrInvalidPassword
	}

	return nil
}

func validateMobileNumber(mobileNumber string) error {
	mobileNumber = strings.TrimSpace(mobileNumber)

	if mobileNumber == "" {
		return nil
	}

	if !phoneRegex.MatchString(mobileNumber) {
		return ErrInvalidMobileNumber
	}

	return nil
}

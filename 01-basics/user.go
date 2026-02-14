package basics

import (
	"errors"
	"regexp"
	"strings"
)

type User struct {
	Username string
	Email    string
	Age      int
}

var (
	ErrEmptyUsername    = errors.New("username cannot be empty")
	ErrInvalidEmail     = errors.New("invalid email format")
	ErrInvalidAge       = errors.New("age must be between 0 and 150")
	ErrUsernameTooShort = errors.New("username must be at least 3 characters")
	ErrUsernameTooLong  = errors.New("username must be less than 30 characters")
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func ValidateUser(u User) error {
	if err := ValidateUsername(u.Username); err != nil {
		return err
	}
	if err := ValidateEmail(u.Email); err != nil {
		return err
	}
	if err := ValidateAge(u.Age); err != nil {
		return err
	}
	return nil
}

func ValidateUsername(username string) error {
	username = strings.TrimSpace(username)

	if username == "" {
		return ErrEmptyUsername
	}

	if len(username) < 3 {
		return ErrUsernameTooShort
	}

	if len(username) > 20 {
		return ErrUsernameTooLong
	}

	return nil
}

func ValidateEmail(email string) error {
	email = strings.TrimSpace(email)

	if email == "" {
		return ErrInvalidEmail
	}

	if !emailRegex.MatchString(email) {
		return ErrInvalidEmail
	}

	return nil
}

func ValidateAge(age int) error {
	if age < 0 || age > 150 {
		return ErrInvalidAge
	}
	return nil
}

func SanitizeUsername(username string) string {
	return strings.ToLower(strings.TrimSpace(username))
}

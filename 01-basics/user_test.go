package basics

import (
	"errors"
	"testing"
)

func TestValidateUsernameValid(t *testing.T) {
	username := "deepanshu_mehra"
	err := ValidateUsername(username)

	if err != nil {
		t.Errorf("ValidateUsername(%q) returned error: %v; want nil", username, err)
	}
}

func TestValidateUsernameEmpty(t *testing.T) {
	username := ""
	err := ValidateUsername(username)

	if err == nil {
		t.Fatal("expected error for empty username, got nil")
	}

	if !errors.Is(err, ErrEmptyUsername) {
		t.Errorf("error = %v; want ErrEmptyUsername", err)
	}
}

func TestValidateUsernameTooShort(t *testing.T) {
	username := "dm"
	err := ValidateUsername(username)

	if err == nil {
		t.Fatal("expected error for too short username, got nil")
	}

	if !errors.Is(err, ErrUsernameTooShort) {
		t.Errorf("error = %v; want ErrUsernameTooShort", err)
	}
}

func TestValidateUsernameTooLong(t *testing.T) {
	username := "deepanshu_mehra_is_long_name"
	err := ValidateUsername(username)

	if err == nil {
		t.Fatal("expected error for too long username, got nil")
	}

	if !errors.Is(err, ErrUsernameTooLong) {
		t.Errorf("error = %v; want ErrUsernameTooLong", err)
	}
}

func TestValidateEmailValid(t *testing.T) {
	email := "john@example.com"
	err := ValidateEmail(email)

	if err != nil {
		t.Errorf("ValidateEmail(%q) returned error: %v; want nil", email, err)
	}
}

func TestValidateEmailInvalid(t *testing.T) {
	invalidEmails := []string{
		"",
		"notanemail",
		"missing@domain",
		"@nodomain.com",
		"no@domain",
	}

	for _, email := range invalidEmails {
		err := ValidateEmail(email)
		if err == nil {
			t.Errorf("ValidateEmail(%q) returned nil; want error", email)
		}
		if !errors.Is(err, ErrInvalidEmail) {
			t.Errorf("ValidateEmail(%q) error = %v; want ErrInvalidEmail", email, err)
		}
	}
}

func TestValidateAgeValid(t *testing.T) {
	validAges := []int{0, 1, 18, 25, 65, 100, 150}

	for _, age := range validAges {
		err := ValidateAge(age)
		if err != nil {
			t.Errorf("ValidateAge(%d) returned error: %v; want nil", age, err)
		}
	}
}

func TestValidateAgeInvalid(t *testing.T) {
	invalidAges := []int{-1, 151, -100, 200}

	for _, age := range invalidAges {
		err := ValidateAge(age)
		if err == nil {
			t.Errorf("ValidateAge(%d) expected error returned nil", age)
		}
		if !errors.Is(err, ErrInvalidAge) {
			t.Errorf("error = %v; want ErrInvalidAge", err)
		}
	}
}

func TestSanitizaUsername(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{" JohnDoe", "johndoe"},
		{"ALICE", "alice"},
		{"bob", "bob"},
		{" MixedCase ", "mixedcase"},
	}

	for _, tt := range tests {
		result := SanitizeUsername(tt.input)
		if result != tt.expected {
			t.Errorf("SanitizeUsername(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

func TestValidateUser(t *testing.T) {
	// Test valid user
	validUser := User{
		Username: "johndoe",
		Email:    "john@example.com",
		Age:      25,
	}

	err := ValidateUser(validUser)
	if err != nil {
		t.Errorf("ValidateUser(validUser) returned error: %v; want nil", err)
	}

	// Invalid Test
	invalidUser1 := User{
		Username: "",
		Email:    "developer@test.com",
		Age:      22,
	}
	err = ValidateUser(invalidUser1)
	if err == nil {
		t.Error("expected error for invalid user go nil")
	}
	if !errors.Is(err, ErrEmptyUsername) {
		t.Errorf("error = %v; want ErrEmptyUsername", err)
	}
}

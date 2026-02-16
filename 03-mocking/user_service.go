package mocking

import (
	"errors"
	"fmt"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidUserData   = errors.New("invalid user data")
	ErrDatabaseError     = errors.New("database error")
	ErrUserAlreadyExists = errors.New("user already exists")
)

type User struct {
	ID       int
	Username string
	Email    string
	Active   bool
}

type UserRepository interface {
	GetByID(id int) (*User, error)
	GetByUsername(username string) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id int) error
}

type EmailSender interface {
	SendWelcomeEmail(email, username string) error
	SendPasswordResetEmail(email, token string) error
}

type UserService struct {
	repo  UserRepository
	email EmailSender
}

func NewUserService(repo UserRepository, email EmailSender) *UserService {
	return &UserService{
		repo:  repo,
		email: email,
	}
}

func (s *UserService) RegisterUser(username, email string) (*User, error) {
	if username == "" || email == "" {
		return nil, ErrInvalidUserData
	}

	existing, err := s.repo.GetByUsername(username)
	if err == nil && existing != nil {
		return nil, ErrUserAlreadyExists
	}

	user := &User{
		Username: username,
		Email:    email,
		Active:   true,
	}

	err = s.repo.Create(user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Send welcome email (don't fail registration if email fails)
	_ = s.email.SendWelcomeEmail(email, username)

	return user, nil
}

func (s *UserService) GetUser(id int) (*User, error) {
	if id <= 0 {
		return nil, ErrInvalidUserData
	}

	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func (s *UserService) DeactivateUser(id int) error {
	user, err := s.GetUser(id)
	if err != nil {
		return err
	}

	user.Active = false
	return s.repo.Update(user)
}

func (s *UserService) ActivateUser(id int) error {
	user, err := s.GetUser(id)
	if err != nil {
		return err
	}

	user.Active = true
	return s.repo.Update(user)
}

func (s *UserService) UpdateEmail(id int, newEmail string) error {
	if newEmail == "" {
		return ErrInvalidUserData
	}

	user, err := s.GetUser(id)
	if err != nil {
		return err
	}

	oldEmail := user.Email
	user.Email = newEmail

	err = s.repo.Update(user)
	if err != nil {
		return err
	}

	// Send notification to old email (don't fail update if email fails)
	_ = s.email.SendWelcomeEmail(oldEmail, fmt.Sprintf("Your email was changed to %s", newEmail))

	return nil
}

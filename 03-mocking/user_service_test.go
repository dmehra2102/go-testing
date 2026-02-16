package mocking

import (
	"errors"
	"testing"
)

type MockUserRepository struct {
	users map[int]*User

	// Track calls for verification
	GetByIDCalls       []int
	GetByUsernameCalls []string
	CreateCalls        []*User
	UpdateCalls        []*User
	DeleteCalls        []int

	// Configure behavior
	GetByIDError       error
	GetByUsernameError error
	CreateError        error
	UpdateError        error
	DeleteError        error

	nextID int
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users:  make(map[int]*User),
		nextID: 1,
	}
}

func (m *MockUserRepository) GetByID(id int) (*User, error) {
	m.GetByIDCalls = append(m.GetByIDCalls, id)

	if m.GetByIDError != nil {
		return nil, m.GetByIDError
	}

	user, exists := m.users[id]
	if !exists {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func (m *MockUserRepository) GetByUsername(username string) (*User, error) {
	m.GetByUsernameCalls = append(m.GetByUsernameCalls, username)

	if m.GetByUsernameError != nil {
		return nil, m.GetByUsernameError
	}

	for _, user := range m.users {
		if user.Username == username {
			return user, nil
		}
	}

	return nil, ErrUserNotFound
}

func (m *MockUserRepository) Create(user *User) error {
	m.CreateCalls = append(m.CreateCalls, user)

	if m.CreateError != nil {
		return m.CreateError
	}

	user.ID = m.nextID
	m.nextID++
	m.users[user.ID] = user

	return nil
}

func (m *MockUserRepository) Update(user *User) error {
	m.UpdateCalls = append(m.UpdateCalls, user)

	if m.UpdateError != nil {
		return m.UpdateError
	}

	if _, exists := m.users[user.ID]; !exists {
		return ErrUserNotFound
	}

	m.users[user.ID] = user
	return nil
}

func (m *MockUserRepository) Delete(id int) error {
	m.DeleteCalls = append(m.DeleteCalls, id)

	if m.DeleteError != nil {
		return m.DeleteError
	}

	delete(m.users, id)
	return nil
}

type MockEmailSender struct {
	WelcomeEmailCalls      []WelcomeEmailCall
	PasswordResetCalls     []PasswordResetCall
	SendWelcomeEmailError  error
	SendPasswordResetError error
}

type WelcomeEmailCall struct {
	Email    string
	Username string
}

type PasswordResetCall struct {
	Email string
	Token string
}

func NewMockEmailSender() *MockEmailSender {
	return &MockEmailSender{}
}

func (m *MockEmailSender) SendWelcomeEmail(email, username string) error {
	m.WelcomeEmailCalls = append(m.WelcomeEmailCalls, WelcomeEmailCall{
		Email:    email,
		Username: username,
	})
	return m.SendWelcomeEmailError
}

func (m *MockEmailSender) SendPasswordResetEmail(email, token string) error {
	m.PasswordResetCalls = append(m.PasswordResetCalls, PasswordResetCall{
		Email: email,
		Token: token,
	})
	return m.SendPasswordResetError
}

func TestRegisterUser(t *testing.T) {
	tests := []struct {
		name          string
		username      string
		email         string
		existingUser  *User
		repoError     error
		expectedError error
		checkEmail    bool
	}{
		{
			name:          "successful registration",
			username:      "johndoe",
			email:         "john@example.com",
			existingUser:  nil,
			repoError:     nil,
			expectedError: nil,
			checkEmail:    true,
		},
		{
			name:          "empty username",
			username:      "",
			email:         "john@example.com",
			expectedError: ErrInvalidUserData,
			checkEmail:    false,
		},
		{
			name:          "empty email",
			username:      "johndoe",
			email:         "",
			expectedError: ErrInvalidUserData,
			checkEmail:    false,
		},
		{
			name:     "username already exists",
			username: "johndoe",
			email:    "john@example.com",
			existingUser: &User{
				ID:       1,
				Username: "johndoe",
				Email:    "existing@example.com",
			},
			expectedError: ErrUserAlreadyExists,
			checkEmail:    false,
		},
		{
			name:          "database error",
			username:      "johndoe",
			email:         "john@example.com",
			repoError:     ErrDatabaseError,
			expectedError: ErrDatabaseError,
			checkEmail:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockUserRepository()
			mockEmail := NewMockEmailSender()

			if tt.existingUser != nil {
				mockRepo.users[tt.existingUser.ID] = tt.existingUser
			}
			if tt.repoError != nil {
				mockRepo.CreateError = tt.repoError
			}

			service := NewUserService(mockRepo, mockEmail)

			// Execute
			user, err := service.RegisterUser(tt.username, tt.email)

			// Verify error
			if tt.expectedError != nil {
				if err == nil {
					t.Fatalf("expected error %v, got nil", tt.expectedError)
				}
				if !errors.Is(err, tt.expectedError) && err.Error() != tt.expectedError.Error() {
					t.Errorf("error = %v; want %v", err, tt.expectedError)
				}
				return
			}

			// Verify success
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if user == nil {
				t.Fatal("expected user, got nil")
			}

			if user.Username != tt.username {
				t.Errorf("username = %q; want %q", user.Username, tt.username)
			}

			if user.Email != tt.email {
				t.Errorf("email = %q; want %q", user.Email, tt.email)
			}

			if !user.Active {
				t.Error("user should be active")
			}

			// Verify email was sent
			if tt.checkEmail {
				if len(mockEmail.WelcomeEmailCalls) != 1 {
					t.Errorf("expected 1 welcome email, got %d", len(mockEmail.WelcomeEmailCalls))
				} else {
					call := mockEmail.WelcomeEmailCalls[0]
					if call.Email != tt.email {
						t.Errorf("welcome email sent to %q; want %q", call.Email, tt.email)
					}
					if call.Username != tt.username {
						t.Errorf("welcome email username %q; want %q", call.Username, tt.username)
					}
				}
			}

		})
	}
}

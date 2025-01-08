package mocks

import (
	"time"

	"github.com/AlessioPani/go-snippetbox/internal/models"
)

var mockUser = models.User{
	ID:             1,
	Name:           "John Doe",
	Email:          "test@test.com",
	HashedPassword: []byte("password"),
	Created:        time.Now(),
}

type UserModel struct{}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "duplicate@mail.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	if email == "test@test.com" && password == "password" {
		return 1, nil
	}
	return 0, models.ErrInvalidCredentials
}

func (m *UserModel) EmailTaken(email string) (bool, error) {
	switch email {
	case "duplicate@mail.com":
		return true, nil
	default:
		return false, nil
	}
}

func (m *UserModel) Exists(id int) (bool, error) {
	switch id {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}

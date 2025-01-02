package models

import (
	"database/sql"
	"time"
)

// User is a struct containing the user data.
type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

// UserModel is a struct used to call DB operations.
type UserModel struct {
	DB *sql.DB
}

// Insert adds a new record to the Users table.
func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

// Authenticate is used to verify whether a user exists with the provided email address and password.
// This will return the relevant user ID if they do.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Exists is used to check if a user exists with a specific ID.
func (m *UserModel) Exists(id int) (bool, error) {
	return false, nil
}

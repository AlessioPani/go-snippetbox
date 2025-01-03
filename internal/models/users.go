package models

import (
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
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
	// Create a bcrypt hash of the plain-text password.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (name, email, hashed_password, created)
 			 VALUES(?, ?, ?, datetime())`

	// Execute the query, populating the placeholders. If errors were found, return it
	_, err = m.DB.Exec(query, name, email, string(hashedPassword))
	if err != nil {
		return err
	}

	return nil
}

// Authenticate is used to verify whether a user exists with the provided email address and password.
// This will return the relevant user ID if they do.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	var id int
	var hashedPassword []byte

	query := "SELECT id, hashed_password FROM users WHERE email = ?"

	err := m.DB.QueryRow(query, email).Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	// Check whether the hashed password and plain-text password provided match.
	// If they don't, we return the ErrInvalidCredentials error.
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	// Otherwise, the password is correct. Return the user ID.
	return id, nil
}

// EmailTaken is used to check if a mail exists already.
func (m *UserModel) EmailTaken(email string) (bool, error) {
	var exists bool

	query := "SELECT EXISTS(SELECT true FROM users WHERE email = ?)"

	// This query is expected to get 1 row at most.
	result := m.DB.QueryRow(query, email)

	// Copy the result into the boolean variable.
	// We need to check only if a row was returned, and the query will return
	// a 1 (true) if at least one row was returned.
	err := result.Scan(&exists)

	return exists, err
}

// Exists is used to check if a user exists with a specific ID.
func (m *UserModel) Exists(id int) (bool, error) {
	return false, nil
}

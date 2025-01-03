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
	return 0, nil
}

// EmailTaken is used to check if a mail exists already.
func (m *UserModel) EmailTaken(email string) (bool, error) {
	query := `SELECT id FROM users WHERE email = ?`

	// This query is expected to get 1 row at most.
	result := m.DB.QueryRow(query, email)

	// Copy the result into a User struct and check for errors.
	// We need to check only if a row was returned, so there is no need to
	// get all the fields.
	var u User
	err := result.Scan(&u.ID)
	if err != nil {
		// Check if Scan didn't return any rows
		// If so, the mail in input can be utilized for a signup.
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil // mail not found
		} else {
			return false, err
		}
	}

	// If no error has occurred, a mail was actually found in the DB.
	return true, nil // mail found
}

// Exists is used to check if a user exists with a specific ID.
func (m *UserModel) Exists(id int) (bool, error) {
	return false, nil
}

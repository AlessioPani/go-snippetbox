package models

import (
	"database/sql"
	"errors"
	"strconv"
	"time"
)

// Snippet is a struct containing the snippet data.
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// SnippetModel is a struct used to call DB operations.
type SnippetModel struct {
	DB *sql.DB
}

// Insert is a function used to insert a snippet on the DB.
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	// Sqlite's datetime('now', <modifier>) doesn't work well with placeholders due to the
	// type of the modifier, which is a composed string, so strconv.Itoa was used as a quick workaround
	query := `INSERT INTO snippets (title, content, created, expires)
			  VALUES(?, ?, datetime(), datetime('now','+` + strconv.Itoa(expires) + " days'))"

	// Execute the query, populating the placeholders. If errors were found, return it
	result, err := m.DB.Exec(query, title, content)
	if err != nil {
		return 0, err
	}

	// Get the last insert id. If errors were found, return it
	// This method is not compatible with all the sql DBs.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get is a method used to get a snippet based on its ID.
func (m *SnippetModel) Get(id int) (Snippet, error) {
	query := `SELECT id, title, content, created, expires FROM snippets
			  WHERE expires > datetime() AND id = ?`

	// Execute the query and store the result (a single row at most) in a *sql.Row type
	result := m.DB.QueryRow(query, id)

	var s Snippet

	// Copy the result into a Snippet struct and check for errors
	err := result.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		// Check if Scan didn't return any rows
		// If so, returns an empty Snippet struct and the custom ErrNoRecord error
		// Otherwise, returns an empty Snippet struct with the received error
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}

	// If Scan ended with no errors, return the filled Snippet struct
	return s, nil
}

// Latest is a method used to get the latest 10 valid snippets.
func (m *SnippetModel) Latest() ([]Snippet, error) {
	query := `SELECT id, title, content, created, expires FROM snippets
			  WHERE expires > datetime() ORDER BY id DESC LIMIT 10`

	results, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	// Close the connection to the DB connection pool after this method
	// returns something.
	defer results.Close()

	var snippets []Snippet

	for results.Next() {
		var s Snippet
		err := results.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	// Check againg for errors after the results iteration.
	if err = results.Err(); err != nil {
		return nil, err
	}

	// If everything went OK, return the snippets slice.
	return snippets, nil
}

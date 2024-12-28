package models

import "errors"

// ErrNoRecord is a custom error which occurs when no rows has been retrieved by a SQL query.
var ErrNoRecord = errors.New("models: no matching record found")

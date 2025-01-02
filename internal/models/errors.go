package models

import "errors"

// ErrNoRecord is a custom error which occurs when no rows
// has been retrieved by a SQL query.
var ErrNoRecord = errors.New("models: no matching record found")

// ErrInvalidCredentials is a custom error which occurs when a user
// tries to login with an incorrect email address or password.
var ErrInvalidCredentials = errors.New("models: invalid credentials")

// ErrDuplicateEmail is a custom error which occurs when a user
// tries to signup with an email address that's already in use.
var ErrDuplicateEmail = errors.New("models: duplicate email")

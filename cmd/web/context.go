package main

type contextKey string

// isAuthenticatedContextKey is the key used to store the isAuthenticated flag.
const isAuthenticatedContextKey = contextKey("isAuthenticated")

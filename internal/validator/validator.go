package validator

import (
	"regexp"
	"slices"
	"strings"
	"unicode/utf8"
)

// EmailRX is a compiled regular expression for checking the basic format of an email address.
// It is compiled once at startup so validation does not recompile it for each request.
var EmailRX = regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)

// Validator is a struct which contains a map of validation error messages.
type Validator struct {
	NonFieldErrors []string
	FieldErrors    map[string]string
}

// Valid returns true if FieldErrors doesn't have any entry in it.
func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0 && len(v.NonFieldErrors) == 0
}

// AddNonFieldError adds a  error messages to the new NonFieldErrors slice.
func (v *Validator) AddNonFieldError(message string) {
	v.NonFieldErrors = append(v.NonFieldErrors, message)
}

// AddFieldError adds an entry in the FieldErrors map.
func (v *Validator) AddFieldError(key string, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}

	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}
}

// CheckField adds an error message if the validation check is not 'ok'.
func (v *Validator) CheckField(ok bool, key string, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

// NotBlank checks if a string is not blank.
func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

// MinChars() returns true if a value contains at least n characters.
func MinChars(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n
}

// MaxChars checks if a value contains no more than n characters.
func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

// PermittedValue checks if a value is contained in a list of permitted values.
func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	return slices.Contains(permittedValues, value)
}

// Matches() returns true if a value matches a provided compiled regular
// expression pattern (e.g: email format).
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

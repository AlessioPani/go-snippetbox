package validator

import (
	"regexp"
	"slices"
	"strings"
	"unicode/utf8"
)

// EmailRX is regular expression pattern for sanity checking the format of an email address.
// This returns a pointer to a 'compiled' regexp.Regexp type, or panics in the event of an error.
// Parsing this pattern once at startup and storing the compiled *regexp.Regexp in a variable
// is more performant than re-parsing the pattern each time we need it.
var EmailRX = regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)

// Validator is a struct which contains a map of validation error messages.
type Validator struct {
	FieldErrors map[string]string
}

// Valid returns true if FieldErrors doesn't have any entry in it.
func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0
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

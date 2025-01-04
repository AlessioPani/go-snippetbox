package main

import (
	"testing"
	"time"

	"github.com/AlessioPani/go-snippetbox/internal/assert"
)

func TestHumanDate(t *testing.T) {
	// Create a slice of anonymous structs containing the test case name,
	// input to our humanDate() function (the tm field), and expected output
	// (the want field).
	tests := []struct {
		name           string
		inputTime      time.Time
		expectedResult string
	}{
		{"UTC", time.Date(2024, 3, 17, 10, 15, 0, 0, time.UTC), "17 Mar 2024 at 10:15"},
		{"Empty", time.Time{}, ""},
		{"CET", time.Date(2024, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)), "17 Mar 2024 at 09:15"},
	}

	// Loop over the test cases.
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := humanDate(test.inputTime)
			assert.Equal(t, test.name, result, test.expectedResult)
		})
	}
}

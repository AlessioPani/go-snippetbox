package main

import (
	"testing"
	"time"

	"github.com/AlessioPani/go-snippetbox/internal/assert"
)

func TestHumanDate(t *testing.T) {
	// Create a slice of anonymous structs containing the test case name,
	// input to our humanDate() function (the inputTime field), and expected output
	// (the expectedResult field).
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
			assert.Equal(t, result, test.expectedResult)
		})
	}
}

func TestAddNumbers(t *testing.T) {
	// Create a slice of anonymous structs containing the test case name,
	// input to our addNumbers() function (the two integers), and expected output
	// (the expectedResult field).
	tests := []struct {
		name           string
		x              int
		y              int
		expectedResult int
	}{
		{"Positive numbers", 1, 3, 4},
		{"Negative numbers", 3, -6, -3},
		{"Nil", 3, -3, 0},
	}

	// Loop over the test cases.
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := addNumbers(test.x, test.y)
			assert.Equal(t, result, test.expectedResult)
		})
	}
}

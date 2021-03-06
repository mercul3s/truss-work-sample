package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	testCases := []struct {
		name          string
		timeString    string
		expectedValue string
		expectedError error
	}{
		{
			name:          "correctly parses timestamp as iso-8601",
			timeString:    "4/1/11 11:00:00 AM",
			expectedValue: "2011-04-01T11:00:00-0500",
			expectedError: nil,
		},
	}
	for _, tc := range testCases {
		result, err := parseTime(tc.timeString)
		assert.Equal(t, tc.expectedValue, result)
		assert.Equal(t, tc.expectedError, err)
	}
}

func TestZip(t *testing.T) {
	testCases := []struct {
		name          string
		zipString     string
		expectedValue string
		expectedError bool
	}{
		{
			name:          "correctly formatted zipcode returns same formatting",
			zipString:     "94121",
			expectedValue: "94121",
			expectedError: false,
		},
		{
			name:          "four digit zipcode returns 0-prefixed five digit zipcode",
			zipString:     "1231",
			expectedValue: "01231",
			expectedError: false,
		},
	}
	for _, tc := range testCases {
		result, err := zip(tc.zipString)
		assert.Equal(t, tc.expectedValue, result)
		compareError(t, tc.expectedError, err)
	}

}

func TestDuration(t *testing.T) {
	testCases := []struct {
		name           string
		durationString string
		expectedValue  float32
		expectedError  bool
	}{
		{
			name:           "parses and returns floating point seconds from HH:MM:SS:MS time string",
			durationString: "1:23:32.123",
			expectedValue:  5012.123,
			expectedError:  false,
		},
		{
			name:           "returns an error for non-time value",
			durationString: "hello",
			expectedValue:  0,
			expectedError:  true,
		},
	}
	for _, tc := range testCases {
		result, err := duration(tc.durationString)
		assert.Equal(t, tc.expectedValue, result)
		compareError(t, tc.expectedError, err)
	}
}

func TestCapitalize(t *testing.T) {
	testCases := []struct {
		name          string
		originalValue string
		expectedValue string
	}{
		{
			name:          "returns string with first letter capitalized",
			originalValue: "hello world",
			expectedValue: "Hello World",
		},
		{
			name:          "returns string with first letter capitalized international",
			originalValue: "Superman übertan",
			expectedValue: "Superman Übertan",
		},
	}

	for _, tc := range testCases {
		result := capitalize(tc.originalValue)
		assert.Equal(t, tc.expectedValue, result)
	}
}

func compareError(t *testing.T, expected bool, err error) {
	if expected && err == nil {
		t.Fatal("expected error, but received nil")
	} else if !expected && err != nil {
		t.Fatalf("error check failed: %s", err.Error())
	}
}

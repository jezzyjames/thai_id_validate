package main

import "testing"

func TestThaiIDValidate(t *testing.T) {
	testcases := []struct {
		given    string
		expected string
	}{
		{
			given:    "1234567890121",
			expected: "nil",
		},
		{
			given:    "123456789012",
			expected: "id digits incorrect",
		},
		{
			given:    "1234567890122",
			expected: "id incorrect",
		},
	}

	for _, testcase := range testcases {
		t.Run("", func(t *testing.T) {
			actual := ValidateThaiID(testcase.given)

			if testcase.expected != actual.Error() {
				t.Errorf("given an ID %s expected error is %v but actual error is %v", testcase.given, testcase.expected, actual)
			}
		})
	}
}

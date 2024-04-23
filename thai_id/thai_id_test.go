package thai_id

import (
	"testing"
)

func TestThaiIDValidate(t *testing.T) {
	testcases := []struct {
		given    string
		expected string
	}{
		{
			given:    "1234567890121",
			expected: "",
		},
		{
			given:    "123456789012",
			expected: "id digits incorrect",
		},
		{
			given:    "1234567890122",
			expected: "id incorrect",
		},
		{
			given:    "1103900018941",
			expected: "",
		},
	}

	for _, testcase := range testcases {
		t.Run("", func(t *testing.T) {
			actual := ValidateThaiID(testcase.given)

			if testcase.expected == "" && actual == nil {
				return
			}

			if testcase.expected != actual.Error() {
				t.Errorf("given an ID %s expected error is %v but actual error is %v", testcase.given, testcase.expected, actual)
			}
		})
	}
}

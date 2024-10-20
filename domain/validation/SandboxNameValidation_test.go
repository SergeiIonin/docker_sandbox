package validation

import (
	"testing"
)

func TestSandboxNameValidation(t *testing.T) {
	snv := NewSandboxNameValidation()

	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Alphanumeric", "sandbox123", "sandbox123"},
		{"With underscore", "sandbox_123", "sandbox_123"},
		{"Ends with underscore", "sandbox_", "sandbox_"},
		{"Has >1 letter", "s_", "s_"},
		{"Starts with number", "1sandbox", ""},
		{"Has one letter", "s", ""},
		{"Starts with _", "_sandbox", ""},
		{"Starts with number and _", "1_sandbox", ""},
	}

	for _, tc := range testCases {
		t.Logf("Running test case: %s\n", tc.name)
		t.Run(tc.name, func(t *testing.T) {
			result, err := snv.ValidateSandboxName(tc.input)
			if result != tc.expected && err != nil {
				t.Errorf("got %v, want %v", result, tc.expected)
			}
		})
	}
}

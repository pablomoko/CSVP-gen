package processors

import (
	"testing"
)

func TestDivisorProcessor_Process(t *testing.T) {
	processor := DivisorProcessor{Divisor: 2}

	tests := []struct {
		input    string
		expected string
	}{
		{"10", "5"},
		{"-10", "-5"},
		{"0", "0"},
		{"abc", "Error: no se pudo convertir el valor a número"},
		{"9", "4"}, // testing integer division rounding down
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := processor.Process(test.input)
			if result != test.expected {
				t.Errorf("Expected %s but got %s", test.expected, result)
			}
		})
	}
}

package main

import "testing"

func TestAdd(t *testing.T) {
	testCases := []struct {
		num1, num2 float64
		expected   float64
	}{
		{2.0, 3.1, 5.1},
		{-2.1, 3.1, 1.0},
		{2.0, -3.0, -1.0},
		{-2.5, -3.5, -6.0},
	}

	for _, tc := range testCases {
		result := Addition(tc.num1, tc.num2)
		if result != tc.expected {
			t.Errorf("Addition(%.2f, %.2f) = %.2f; expected %.2f", tc.num1, tc.num2, result, tc.expected)
		}
	}
}

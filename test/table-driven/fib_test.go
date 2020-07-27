package main

import "testing"

func TestFib(t *testing.T) {
	var fibTests = []struct {
		in       int // input
		expected int // expected result
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
	}

	for _, tt := range fibTests {
		actual := Fib(tt.in)
		if actual != tt.expected {
			t.Errorf("Fib(%d) = %d; expected %d", tt.in, actual, tt.expected)
		}
	}
}

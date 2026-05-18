package main

import "testing"

func TestCalculation(t *testing.T) {

	char := 'A'
	expectedStartLine := (int(char-32) * 9) + 1
	if expectedStartLine != 298 {
		t.Errorf("expected line 298 of 'A', got %d", expectedStartLine)

	}

}

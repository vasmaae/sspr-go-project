package main

import "testing"

func TestGetAverage(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	expected := float32(3.5)
	result := getAverage(&matrix)

	if result != expected {
		t.Errorf("expected %f, got %f", expected, result)
	}
}

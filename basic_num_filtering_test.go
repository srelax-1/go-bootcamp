package main

import (
	"testing"
)

func TestEvenNums(t *testing.T) {
	result := EvenNums([]int{1,2,3,4,5,6})

	if result == nil {
		t.Error("This is an empty array")
	}

	for _, val := range result {
		if val % 2 != 0 {
			t.Errorf("Expected 0 got %d", val)
		}
	}
}

func TestOddNums(t *testing.T) {
	result := OddNums([]int{1,2,3,4,5,6})
	
	if result == nil {
		t.Error("This is an empty array")
	}

	for _, val := range result {
		if val % 2 != 1 {
			t.Errorf("Expected 0 got %d", val)
		}
	}
}
package main

import (
	"testing"
	"slices"
)

func TestEvenNums(t *testing.T) {
	result := EvenNums([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	if result == nil {
		t.Fatal("This is an empty array")
	}

	for _, val := range result {
		if val % 2 != 0 {
			t.Errorf("Expected 0 got %d", val)
		}
	}

	expected := []int{2, 4, 6, 8, 10}
	if !slices.Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestOddNums(t *testing.T) {
	result := OddNums([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	
	if result == nil {
		t.Fatal("This is an empty array")
	}

	for _, val := range result {
		if val % 2 != 1 {
			t.Errorf("Expected 0 got %d", val)
		}
	}

	expected := []int{1, 3, 5, 7, 9}
	if !slices.Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPrimeNums(t *testing.T) {
	result := PrimeNums([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	slices.Sort(result)

	if len(result) == 0 {
		t.Fatal("Expected prime numbers, got an empty array")
	}

	for _, num := range result {
		if num <= 1 {
			t.Errorf("%d is not a prime number", num)
            continue
        }

		// This keeps dividing the number until i gets to the square root of the number being divided.
		for i := 2; i*i <= num; i++ {
			if num % i == 0 {
				t.Errorf("%d is not a prime number", num)
				break
			}
    	}
	}

	expected := []int{2, 3, 5, 7}
	if !slices.Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPrimeOddNums(t *testing.T) {
    result := PrimeOddNums([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
    slices.Sort(result)

    if len(result) == 0 {
        t.Fatal("Expected odd prime numbers, got an empty array")
    }

    for _, num := range result {
        if num <= 1 {
            t.Errorf("%d is not a prime number", num)
            continue
        }

        if num%2 == 0 {
            t.Errorf("%d is an even number", num)
            continue
        }

        for i := 2; i*i <= num; i++ {
            if num%i == 0 {
                t.Errorf("%d is not a prime number", num)
                break
            }
        }
    }

    expected := []int{3, 5, 7}

    if !slices.Equal(result, expected) {
        t.Errorf("Expected %v, got %v", expected, result)
    }
}
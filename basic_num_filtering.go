package main

import "fmt"

// Even Number filtering
func EvenNums(nums []int) []int{
	var result []int
	for _, val := range nums {
		if val % 2 == 0 {
			result = append(result, val)
		}
	}
	return result
}

// Odd Number filtering
func OddNums(nums []int) []int{
	var result []int
	for _, val := range nums {
		if val % 2 == 1 {
			result = append(result, val)
		}
	}
	return result
}

// Prime Number filtering
func PrimeNums(nums []int) []int{
	var result []int
	for _, val := range nums {
		if val <= 1 {
            continue
        }

		isPrime := true

        for i := 2; i*i <= val; i++ {
            if val%i == 0 {
                isPrime = false
                break
            }
        }

        if isPrime {
            result = append(result, val)
        }
	}
	return result
}

// Prime Odd Number filtering
func PrimeOddNums(nums []int) []int{
	var result []int
	for _, val := range nums {
		if val <= 1 {
            continue
        }

		if val % 2 == 0 {
			continue
		}
		
		isPrime := true

        for i := 2; i*i <= val; i++ {
            if val % i == 0 {
                isPrime = false
                break
            }
        }

        if isPrime {
            result = append(result, val)
        }
	}
	return result
}

func EvenMultiplyFive(nums []int) []int{
	var result []int
	for _, val := range nums {
		if val % 5 == 0 && val % 2 == 0 {
			result = append(result, val)
		}
	}
	return result
}

func main() {
	number := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	// even := EvenNums(number)
	// odd := OddNums(number)
	// prime := PrimeNums(number)
	// primeOdd := PrimeOddNums(number)
	evenmultiplyfive := EvenMultiplyFive(number)
	// fmt.Printf("These are the even number: %d\n", even)
	// fmt.Printf("These are the odd number: %d\n", odd)
	// fmt.Printf("These are the prime number: %d\n", prime)
	// fmt.Printf("These are the prime odd number: %d\n", primeOdd)
	fmt.Printf("These are the Even muliply of five: %d\n", evenmultiplyfive)
}



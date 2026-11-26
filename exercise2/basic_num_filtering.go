package main

import "fmt"

// Story 1: Given a list of integers, write a program to return only the even numbers from this list.
func EvenNums(nums []int) []int{
	var result []int
	for _, val := range nums {
		if val % 2 == 0 {
			result = append(result, val)
		}
	}
	return result
}

// Story 2: Given a list of integers, write a program to return only the odd numbers from this list.
func OddNums(nums []int) []int{
	var result []int
	for _, val := range nums {
		if val % 2 == 1 {
			result = append(result, val)
		}
	}
	return result
}

// Story 3: Given a list of integers, write a program to return only the prime numbers from this list.
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

// Story 4: Given a list of integers, write a program to return only the odd prime numbers from this list.
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

// Story 5: Given a list of integers, write a program to return only the even and multiples of 5 from this list.
func EvenMultiplyOfFive(nums []int) []int{
	var result []int
	for _, val := range nums {
		if val % 5 == 0 && val % 2 == 0 {
			result = append(result, val)
		}
	}
	return result
}

// Story 6: Given a list of integers, write a program to return only the odd and multiples of 3 greater than 10 from this list.
func OddMultiplesOfThreeOverTen(nums []int) []int{
	var result []int
	for _, val := range nums {
		if val > 10 {
			if val % 3 == 0 && val % 2 == 1 {
				result = append(result, val)
			}
		}
	}
	return result
}
// Story 7: Given a list of integers, and a set of conditions (odd, even, greater than 5, multiple of 3, prime, and many more such 
// custom conditions that may be dynamically defined by user), write a program to return only the integers from the given list that 
// match ALL the conditions.

// Story 8: Given a list of integers, and a set of conditions (odd, even, greater than 5, multiple of 3, prime, and many more such 
// custom conditions that may be dynamically defined by user), write a program to return only the integers from the given list that 
// match ANY of the conditions.

func main() {
	number := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	// even := EvenNums(number)
	// odd := OddNums(number)
	// prime := PrimeNums(number)
	// primeOdd := PrimeOddNums(number)
	// evenmultiplyoffive := EvenMultiplyOfFive(number)
	oddmultiplesofthreeoverten := OddMultiplesOfThreeOverTen(number)

	// fmt.Printf("These are the even number: %d\n", even)
	// fmt.Printf("These are the odd number: %d\n", odd)
	// fmt.Printf("These are the prime number: %d\n", prime)
	// fmt.Printf("These are the prime odd number: %d\n", primeOdd)
	// fmt.Printf("These are the Even muliply of five: %d\n", evenmultiplyoffive)
	fmt.Printf("These are the odd numbers that are muliply of three greater than 10: %d\n", oddmultiplesofthreeoverten)
}



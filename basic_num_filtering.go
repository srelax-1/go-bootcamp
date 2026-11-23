package main

// import "fmt"

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

// func main() {
// 	okaya := EvenNums([]int{1,2,3,4,5,6})
// 	fmt.Println(okaya)
// }



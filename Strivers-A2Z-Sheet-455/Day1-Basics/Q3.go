package main

import "fmt"

func main() {
	// Check if 28 is a perfect number
	fmt.Println(checkPerfectNumber(7)) // Output: true
}

// Function to check if a number is a perfect number
func checkPerfectNumber(num int) bool {
	// Base case: numbers less than 2 cannot be perfect numbers
	if num < 2 {
		return false
	}

	allDivisors := findDivisors(num)
	sumOfDivisors := 0

	// Sum all divisors except the number itself
	for _, value := range allDivisors {
		sumOfDivisors += value
	}

	return sumOfDivisors == num
}

// Function to find all proper divisors of a given number
func findDivisors(num int) []int {
	var divisors []int

	// Iterate from 1 to num/2 to find all proper divisors
	for i := 1; i <= num/2; i++ {
		if num%i == 0 {
			divisors = append(divisors, i)
		}
	}

	return divisors
}

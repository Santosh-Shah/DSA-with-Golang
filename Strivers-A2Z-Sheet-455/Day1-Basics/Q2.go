package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse(-4560))
	fmt.Println(reverse(123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(0))
}

func reverse(x int) int {
	// Constants for 32-bit signed integer range
	const (
		INT32_MIN = -1 << 31
		INT32_MAX = (1 << 31) - 1
	)

	reversedNum := 0

	for x != 0 {
		// Get the last digit of x
		lastDigit := x % 10

		// Check for integer overflow before updating reversedNum
		if reversedNum > INT32_MAX/10 || (reversedNum == INT32_MAX/10 && lastDigit > 7) {
			return 0
		}
		if reversedNum < INT32_MIN/10 || (reversedNum == INT32_MIN/10 && lastDigit < -8) {
			return 0
		}

		reversedNum = (reversedNum * 10) + lastDigit
		x = x / 10
	}

	return reversedNum
}

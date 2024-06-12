package main

import "fmt"

func main() {

	//TODO: 2520. Count the Digits That Divide a Number

	fmt.Println("Numbers Count: ", countDigits(1248))

	//temp := 72
	//fmt.Println(temp % 10)
	//fmt.Println(temp / 10)
	//temp = temp / 10
	//fmt.Println(temp)
}

func countDigits(num int) int {
	// Base case
	if num <= 0 {
		return -1
	}

	count := 0
	temp := num // Store the original number

	for num > 0 {
		lastValue := num % 10
		if lastValue != 0 && temp%lastValue == 0 {
			count++
		}
		num = num / 10
	}

	return count
}

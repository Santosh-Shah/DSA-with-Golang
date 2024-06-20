package main

import "fmt"

func main() {
	arr := []int{29, 10, 14, 37, 14}

	fmt.Println("Before Sorting: ")
	for _, value := range arr {
		print(value, " ")
	}

	fmt.Println("\nAfter sorting: ")
	output := bubbleSort(arr)
	for _, value := range output {
		print(value, " ")
	}
}

func bubbleSort(nums []int) []int {
	// Base case
	if len(nums) < 2 {
		return nums
	}

	length := len(nums)

	for i := length - 1; i >= 0; i-- {
		for j := 0; j <= i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

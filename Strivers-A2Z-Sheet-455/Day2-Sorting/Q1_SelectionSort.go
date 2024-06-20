package main

import "fmt"

func main() {
	arr := []int{29, 10, 14, 37, 14}

	fmt.Println("Before Sorting: ")
	for _, value := range arr {
		print(value, " ")
	}

	fmt.Println("\nAfter sorting: ")
	output := selectionSort(arr)
	for _, value := range output {
		print(value, " ")
	}
}

func selectionSort(nums []int) []int {
	// Base case
	if len(nums) < 2 {
		return nums
	}

	for i := 0; i < len(nums)-1; i++ {
		miniValue := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[miniValue] {
				miniValue = j
			}
		}

		nums[i], nums[miniValue] = nums[miniValue], nums[i]
	}
	return nums
}

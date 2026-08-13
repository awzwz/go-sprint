package main

import "fmt"

func sum(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func max(nums []int) int {
	result := nums[0]
	for _, v := range nums {
		if v > result {
			result = v
		}
	}
	return result
}

func reverse(nums []int) []int {
	result := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		result = append(result, nums[i])
	}
	return result
}

func main() {
	nums := []int{5, 2, 9, 1, 7}
	fmt.Println(sum(nums))
	fmt.Println(max(nums))
	fmt.Println(reverse(nums))
}

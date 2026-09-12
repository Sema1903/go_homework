package main

import "fmt"

func twoSum(nums []int, target int) []int {
	result := []int{}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = []int{i, j}
			}
		}
	}
	return result
}
func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4}, 7))
}

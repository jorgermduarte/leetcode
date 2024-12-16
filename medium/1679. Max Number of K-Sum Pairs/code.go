package main

import "fmt"

/*
*
* You are given an integer array nums and an integer k.
* In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.
* Return the maximum number of operations you can perform on the array.
 */

func main() {
	totalOp := maxOperations([]int{1, 2, 3, 4}, 5)
	fmt.Println(totalOp)
}

func maxOperations(nums []int, k int) int {
	totalOperations := 0
	i := 0

	for i < len(nums) {
		current := nums[i]
		if current >= k {
			i++
			continue
		}

		found := false
		for lastIndex := len(nums) - 1; lastIndex >= 0; lastIndex-- {
			if i >= lastIndex {
				break
			}

			second := nums[lastIndex]
			sum := (current + second)
			if sum == k {
				nums = append(nums[:lastIndex], nums[lastIndex+1:]...)
				nums = append(nums[:i], nums[i+1:]...)
				totalOperations++
				found = true
				break
			}
		}
		if !found {
			i++
		}
	}
	return totalOperations
}

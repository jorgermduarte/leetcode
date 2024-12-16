package main

import "fmt"

/*
Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.

Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
Output: 6
Explanation: [1,1,1,0,0,1,1,1,1,1,1]


Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
Output: 10
Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
*/

func main() {
	result := longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3)
	fmt.Println(result)
}
func longestOnes(nums []int, k int) int {
	start, maxLen, zerosCount := 0, 0, 0

	for end := 0; end < len(nums); end++ {
		if nums[end] == 0 {
			zerosCount++
		}

		for zerosCount > k {
			if nums[start] == 0 {
				zerosCount--
			}
			start++
		}

		currentLen := end - start + 1
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}

	return maxLen
}

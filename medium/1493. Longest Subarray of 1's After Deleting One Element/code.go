package main

import "fmt"

/*
Given a binary array nums, you should delete one element from it.

Return the size of the longest non-empty subarray containing only 1's in the resulting array. Return 0 if there is no such subarray.

Example 1:

Input: nums = [1,1,0,1]
Output: 3
Explanation: After deleting the number in position 2, [1,1,1] contains 3 numbers with value of 1's.
Example 2:

Input: nums = [0,1,1,1,0,1,1,0,1]
Output: 5
Explanation: After deleting the number in position 4, [0,1,1,1,1,1,0,1] longest subarray with value of 1's is [1,1,1,1,1].
Example 3:

Input: nums = [1,1,1]
Output: 2
Explanation: You must delete one element.

*/
func main() {
	result := longestSubarray([]int{1, 1, 0, 1})
	fmt.Println(result)
}

func longestSubarray(nums []int) int {
	max := 0
	start := 0
	zeroUsed := 0

	for end := 0; end < len(nums); end++ {
		if nums[end] == 0 {
			zeroUsed++
		}

		for zeroUsed > 1 {
			if nums[start] == 0 {
				zeroUsed--
			}
			start++
		}

		currentLength := end - start + 1
		if currentLength > max {
			max = currentLength
		}
	}

	return max - 1
}

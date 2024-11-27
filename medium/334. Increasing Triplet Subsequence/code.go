/*
Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.

Example 1:

Input: nums = [1,2,3,4,5]
Output: true
Explanation: Any triplet where i < j < k is valid.
Example 2:

Input: nums = [5,4,3,2,1]
Output: false
Explanation: No triplet exists.
Example 3:

Input: nums = [2,1,5,0,4,6]
Output: true
Explanation: The triplet (3, 4, 5) is valid because nums[3] == 0 < nums[4] == 4 < nums[5] == 6.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{2, 1, 5, 0, 4, 6}
	result := increasingTripletWrong(nums)
	fmt.Println(result)
}

func increasingTripletWrong(nums []int) bool {
	output := false
	for i := 1; i < len(nums)-1; i++ {
		if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
			output = true
			break
		}
	}
	return output
}

func increasingTriplet(nums []int) bool {
	m1, m2 := math.MaxInt64, math.MaxInt64
	for _, v := range nums {
		if v <= m1 {
			m1 = v
		} else if v <= m2 {
			m2 = v
		} else {
			return true
		}
	}
	return false
}

/*
238. Product of Array Except Self
Medium
Topics
Companies
Hint
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.



Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]
Example 2:

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
*/

package main

import "fmt"

func main() {
	//nums := []int{-1, 1, 0, -3, 3}
	nums := []int{1, 2, 3, 4}
	result := productExceptSelf(nums)
	fmt.Println(result)
}

// O(n^2)
func productExceptSelf2(nums []int) []int {
	result := make([]int, len(nums))
	for i := range result {
		result[i] = 1
	}

	k := 0
	for i := 0; i < len(nums); {
		if k >= len(nums) {
			break
		}
		if i != k {
			result[k] = result[k] * nums[i]
			// avoid future calculations since the result would be always 0
			if result[k] == 0 {
				k++
				i = 0
				continue
			}
		}
		i++
		if i == len(nums) {
			k++
			i = 0
		}
	}
	return result
}

// O(n)
func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))
	prefixArr := make([]int, len(nums))
	postfixArr := make([]int, len(nums))

	prefix := 1
	for i := 0; i < len(nums); {
		prefixArr[i] = prefix
		prefix = prefix * nums[i]
		i++
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; {
		postfixArr[i] = postfix
		postfix = postfix * nums[i]
		i--
	}

	for i := 0; i < len(nums); i++ {
		output[i] = prefixArr[i] * postfixArr[i]
	}

	return output
}

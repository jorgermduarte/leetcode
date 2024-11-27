/*
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.

Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
Example 2:

Input: nums = [0]
Output: [0]
*/

package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	fmt.Println(nums)
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	output := []int{}
	zerosCount := 0

	for _, v := range nums {
		if v != 0 {
			output = append(output, v)
		} else {
			zerosCount++
		}
	}

	for zerosCount > 0 {
		output = append(output, 0)
		zerosCount--
	}

	// update directly the original slice
	copy(nums, output)
}

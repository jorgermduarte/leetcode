package main

import (
	"fmt"
)

/*
You are given an integer array nums consisting of n elements, and an integer k.
Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.
*/
func main() {
	result := findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4)
	fmt.Println(result)
}

// time limit exceeded
func findMaxAverageBad(nums []int, k int) float64 {

	m := make(map[int]float64)
	for i := 0; i < len(nums); i++ {
		if (i + k) > len(nums) {
			break
		}
		for x := 0; x < (k); x++ {
			current := nums[i+x]
			m[i] += float64(current)
		}
	}

	biggest := m[0]
	for _, v := range m {
		if v > biggest {
			biggest = v
		}
	}

	return float64(biggest / float64(k))
}

func findMaxAverage(nums []int, k int) float64 {
	currentSum := 0
	for i := 0; i < k; i++ {
		currentSum += nums[i]
	}
	maxSum := currentSum
	for i := k; i < len(nums); i++ {
		currentSum = currentSum - nums[i-k] + nums[i]
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return float64(maxSum) / float64(k)
}

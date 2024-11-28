/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.

Example 1:

Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
Example 2:

Input: height = [1,1]
Output: 1
*/
package main

import "fmt"

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result := maxArea(nums)
	fmt.Println(result)
}

// O(n^2)
func maxAreaOn2(height []int) int {
	if len(height) == 0 {
		return 0
	} else if len(height) == 1 {
		return height[0]
	}

	biggestHeight := 0
	for i := 0; i < len(height); i++ {
		for x := len(height) - 1; x > 0; x-- {
			smallestHeight := height[i]

			if height[x] < smallestHeight {
				smallestHeight = height[x]
			}

			width := x - i
			area := smallestHeight * width

			if area > biggestHeight {
				biggestHeight = area
			}
		}
	}

	return biggestHeight
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		h := min(height[left], height[right])
		width := right - left
		maxArea = max(maxArea, h*width)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

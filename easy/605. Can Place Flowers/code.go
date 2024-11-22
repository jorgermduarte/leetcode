/*
You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.

Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n, return true if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and false otherwise.



Example 1:

Input: flowerbed = [1,0,0,0,1], n = 1
Output: true
Example 2:

Input: flowerbed = [1,0,0,0,1], n = 2
Output: false
*/

package main

import "fmt"

func main() {
	arr := []int{1, 0, 0, 0, 1}
	n := 1
	fmt.Print(" can place flowers: ", canPlaceFlowers(arr, n))
}

// n is the number of flowers to be planted
// we cannot have two plants together !
func canPlaceFlowers(flowerbed []int, n int) bool {
	foundSlots := 0

	for flowerIndex := 0; flowerIndex < len(flowerbed); flowerIndex++ {
		fmt.Println("index", flowerIndex)

		// verify if the current slot is available
		if flowerbed[flowerIndex] != 0 {
			continue
		}

		// check prev slot that must be empty
		if (flowerIndex-1) >= 0 && (flowerbed[flowerIndex-1] == 1) {
			continue
		}

		// check next slot that must be empty
		if flowerIndex+1 < len(flowerbed) && (flowerbed[flowerIndex+1] == 1) {
			continue
		}

		foundSlots++
		// change flowerbed at index flowerIndex to refresh changes for the next verificaton
		flowerbed[flowerIndex] = 1
		fmt.Println(flowerbed)

		if foundSlots >= n {
			break
		}
	}

	return foundSlots >= n
}

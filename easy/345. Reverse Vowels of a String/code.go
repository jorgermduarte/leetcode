/*
Given a string s, reverse only all the vowels in the string and return it.

The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.


Example 1:

Input: s = "IceCreAm"
Output: "AceCreIm"

Explanation:

The vowels in s are ['I', 'e', 'e', 'A']. On reversing the vowels, s becomes "AceCreIm".

Example 2:

Input: s = "leetcode"
Output: "leotcede"
*/

package main

import "fmt"

func main() {
	var s string = " apG0i4maAs::sA0m4i0Gp0"
	fmt.Println(s)
	fmt.Println(reverseVowels(s)) // AceCreIm
}

func isVolwel(char byte) bool {
	char = toLower(char)
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'
}

func toLower(char byte) byte {
	if char >= 'A' && char <= 'Z' {
		return char + 32
	}
	return char
}

func reverseVowels(s string) string {
	runes := []rune(s)
	leftIndex, rightIndex := 0, len(runes)-1

	for leftIndex < rightIndex {
		// Find next vowel from the left
		for leftIndex < rightIndex && !isVolwel(byte(runes[leftIndex])) {
			leftIndex++
		}
		// Find next vowel from the right
		for leftIndex < rightIndex && !isVolwel(byte(runes[rightIndex])) {
			rightIndex--
		}

		// Swap vowels and move pointers inward
		if leftIndex < rightIndex {
			runes[leftIndex], runes[rightIndex] = runes[rightIndex], runes[leftIndex]
			leftIndex++
			rightIndex--
		}
	}

	return string(runes)
}

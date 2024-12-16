package main

import "fmt"

/*
Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.
Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.

Input: s = "abciiidef", k = 3
Output: 3

Input: s = "aeiou", k = 2
Output: 2

Input: s = "leetcode", k = 3
Output: 2
*/
func main() {
	result := maxVowels("tnfazcwrryitgacaabwm", 4)
	fmt.Println(result)
}

func maxVowels(s string, k int) int {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	max, currentVowels := 0, 0

	for i := 0; i < k; i++ {
		if vowels[s[i]] {
			currentVowels++
		}
	}
	max = currentVowels

	for i := k; i < len(s); i++ {
		if vowels[s[i]] {
			currentVowels++
		}

		// Remove the character that's sliding out
		if vowels[s[i-k]] {
			currentVowels--
		}

		if currentVowels > max {
			max = currentVowels
		}

	}

	return max
}

func maxVowelsSlow(s string, k int) int {
	max := 0

	for x := 0; x < len(s); x++ {
		if x+k > len(s) {
			break
		}

		totalVowels := 0
		for i := 0; i < k; i++ {
			letter := s[x+i]
			if isVowel(letter) {
				totalVowels++
			}
		}

		if totalVowels > max {
			max = totalVowels
		}
	}
	return max
}

func isVowel(s byte) bool {
	switch s {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}

/*
Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

Example 1:
Input: s = "abc", t = "ahbgdc"
Output: true

Example 2:
Input: s = "axc", t = "ahbgdc"
Output: false
*/

package main

func main() {

}

// Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
// A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).
func isSubsequence(s string, t string) bool {
	// Check if each character of s appears in t in the same order
	position := 0
	for i := 0; i < len(s); i++ {
		char := s[i]
		found := false
		for x := 0; x < len(t); x++ {
			if x <= position && i > 0 {
				continue
			}
			targetchar := t[x]
			if char == targetchar {
				found = true
				position = x
				break
			}
		}
		if found == false {
			return false
		} else {
			continue
		}
	}
	return true
}

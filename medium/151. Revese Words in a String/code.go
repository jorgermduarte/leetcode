/*
Given an input string s, reverse the order of the words.


A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.

Return a string of the words in reverse order concatenated by a single space.

Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.



Example 1:

Input: s = "the sky is blue"
Output: "blue is sky the"
Example 2:

Input: s = "  hello world  "
Output: "world hello"
Explanation: Your reversed string should not contain leading or trailing spaces.
Example 3:

Input: s = "a good   example"
Output: "example good a"
Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.
*/

package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	var s string = "the sky is blue"
	fmt.Println(s)
	fmt.Println(reverseWords(s)) // blue is sky the
}

func reverseWords(s string) string {
	splitted := strings.Split(s, " ")
	result := ""
	for i := len(splitted) - 1; i >= 0; i-- {
		current := splitted[i]
		if current != "" {
			result += current
			if i != 0 {
				result += " "
			}
		}
	}

	return strings.Trim(result, " ")
}

func reverseWords_example2(s string) string {
	fields := strings.Fields(s)
	slices.Reverse(fields)
	return strings.Join(fields, " ")
	for i := len(splitted) - 1; i >= 0; i-- {
		current := splitted[i]
		if current != "" {
			result += current
			if i != 0 {
				result += " "
			}
		}
	}

	return strings.Trim(result, " ")
}

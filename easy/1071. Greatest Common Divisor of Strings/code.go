/*

For two strings s and t, we say "t divides s" if and only if s = t + t + t + ... + t + t (i.e., t is concatenated with itself one or more times).

Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.


Example 1:

Input: str1 = "ABCABC", str2 = "ABC"
Output: "ABC"
Example 2:

Input: str1 = "ABABAB", str2 = "ABAB"
Output: "AB"
Example 3:

Input: str1 = "LEET", str2 = "CODE"
Output: ""
*/

package main

import (
	"fmt"
)

func main() {
	// Test cases
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))  // Output: "ABC"
	fmt.Println(gcdOfStrings("ABABAB", "ABAB")) // Output: "AB"
	fmt.Println(gcdOfStrings("LEET", "CODE"))   // Output: ""
}

// gcdOfStrings finds the largest string x such that x divides both str1 and str2
func gcdOfStrings(str1 string, str2 string) string {
	// If str1 + str2 is not equal to str2 + str1, there is no common divisor
	if str1+str2 != str2+str1 {
		return ""
	}

	// Find the GCD of the lengths of the two strings
	gcdLength := gcd(len(str1), len(str2))

	// The common divisor string will be the prefix of str1 with length gcdLength
	return str1[:gcdLength]
}

// gcd computes the greatest common divisor of two integers
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

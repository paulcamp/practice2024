package main

import (
	"practice2024/basics"
	"strings"
)

/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.


Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.
*/

func IsPalindrome(s string) bool {
	lower := strings.ToLower(s)
	alphaOnly := basics.StripNonAlphaNumericAndSpaces(lower)

	if alphaOnly == "" {
		return true
	}

	return palindrome(alphaOnly)

}

// func palindromev2(str string) bool {
// 	lastIdx := len(str) - 1

// 	for i := 0; i < lastIdx/2 && i < (lastIdx-i); i++ {
// 		if str[i] != str[lastIdx-i] {
// 			return false
// 		}
// 	}
// 	return true
// }

func palindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

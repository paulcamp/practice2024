package main

//https://leetcode.com/problems/minimum-window-substring/description/

/* Given two strings s and t of lengths m and n respectively, return the minimum window
substring of s such that every character in t (including duplicates) is included in the window.

If there is no such substring, return the empty string "".

Example 1:

Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
Example 2:

Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.
Example 3:

Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.
*/

func MinimumWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	// Dictionary to keep a count of all the unique characters in t.
	dictT := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		dictT[t[i]]++
	}

	// Number of unique characters in t that need to be present in the desired window.
	required := len(dictT)

	// Left and Right pointer
	l, r := 0, 0

	// Formed is used to keep track of how many unique characters in t are present in the current window in its desired frequency.
	formed := 0

	// Dictionary to keep a count of all the unique characters in the current window.
	windowCounts := make(map[byte]int)

	// ans tuple of the form (window length, left, right)
	ans := []int{-1, 0, 0}

	for r < len(s) {
		// Add one character from the right to the window
		c := s[r]
		windowCounts[c]++

		// If the frequency of the current character added equals to the desired count in t then increment the formed count by 1.
		if dictT[c] > 0 && windowCounts[c] == dictT[c] {
			formed++
		}

		// Try and contract the window till the point where it ceases to be 'desirable'.
		for l <= r && formed == required {
			c = s[l]

			// Save the smallest window until now.
			if ans[0] == -1 || r-l+1 < ans[0] {
				ans[0] = r - l + 1
				ans[1] = l
				ans[2] = r
			}

			// The character at the position pointed by the `left` pointer is no longer a part of the window.
			windowCounts[c]--
			if dictT[c] > 0 && windowCounts[c] < dictT[c] {
				formed--
			}

			// Move the left pointer ahead, this would help to look for a new window.
			l++
		}

		// Keep expanding the window once we are done contracting.
		r++
	}

	if ans[0] == -1 {
		return ""
	}
	return s[ans[1] : ans[2]+1]
}

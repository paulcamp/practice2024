package main

/*
https://leetcode.com/problems/integer-to-roman/description/


Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

Roman numerals are formed by appending the conversions of decimal place values from highest to lowest.
Converting a decimal place value into a Roman numeral has the following rules:

If the value does not start with 4 or 9, select the symbol of the maximal value that can be subtracted from the input,
 append that symbol to the result, subtract its value, and convert the remainder to a Roman numeral.

 If the value starts with 4 or 9 use the subtractive form representing one symbol subtracted from the following symbol,
 for example, 4 is 1 (I) less than 5 (V): IV and 9 is 1 (I) less than 10 (X): IX.
 Only the following subtractive forms are used: 4 (IV), 9 (IX), 40 (XL), 90 (XC), 400 (CD) and 900 (CM).

Only powers of 10 (I, X, C, M) can be appended consecutively at most 3 times to represent multiples of 10.
You cannot append 5 (V), 50 (L), or 500 (D) multiple times.
If you need to append a symbol 4 times use the subtractive form.

Given an integer, convert it to a Roman numeral.
Example 1:
Input: num = 3749
Output: "MMMDCCXLIX"

Explanation:

3000 = MMM as 1000 (M) + 1000 (M) + 1000 (M)
 700 = DCC as 500 (D) + 100 (C) + 100 (C)
  40 = XL as 10 (X) less of 50 (L)
   9 = IX as 1 (I) less of 10 (X)
Note: 49 is not 1 (I) less of 50 (L) because the conversion is based on decimal places

Example 2:
Input: num = 58
Output: "LVIII"
Explanation:
50 = L
 8 = VIII

Example 3:
Input: num = 1994
Output: "MCMXCIV"

Explanation:
1000 = M
 900 = CM
  90 = XC
   4 = IV


*/

// mine
// method = keep subtracting from target as we build the string.
// slices hold all legitimate combinations
// ignore rules about X only allowed 3 times, its just noise
func ReverseRoman(num int) string {
	numbers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	numerals := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""

	var i = 0

	for i <= len(numbers) && num > 0 {
		if num >= numbers[i] {
			roman += numerals[i]
			num -= numbers[i]
			continue //could use an else but more readable to show multiples
		}

		i += 1

	}

	return roman

}

// func ReverseRomanFail(num int) string {

// 	reverseMap := map[int]string{
// 		1000: "M",
// 		900:  "CM",
// 		500:  "D",
// 		400:  "CD",
// 		100:  "C",
// 		90:   "XC",
// 		50:   "L",
// 		40:   "XL",
// 		10:   "X",
// 		9:    "IX",
// 		5:    "V",
// 		4:    "IV",
// 		1:    "I",
// 	}

// 	roman := ""

// 	for num > 0 {
// 		//FAIL this wont work because range over a MAP is NOT ORDERED !!!
// 		for k, v := range reverseMap {
// 			if num >= k {
// 				roman += v
// 				num -= k
// 			}

// 		}

// 	}

// 	return roman
// }

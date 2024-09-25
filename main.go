package main

import (
	"fmt"
	"practice2024/basics"
)

func main() {
	FizzBuzz()

	countOfPrimes := CountPrimesBasic(100)
	fmt.Printf("\nCount Of Primes Basic up to 100 is [%d]", countOfPrimes)
	countOfPrimes2 := CountPrimesOptimised(100)
	fmt.Printf("\nCount Of Primes Optimized up to 100 is [%d]", countOfPrimes2)

	palinDromeCheck1 := IsPalindrome("A man, a plan, a canal: Panama")
	fmt.Printf("\nExplanation: is 'amanaplanacanalpanama' a palindrome? %v", palinDromeCheck1)
	palinDromeCheck2 := IsPalindrome("raceacar")
	fmt.Printf("\nExplanation: is 'raceacar' a palindrome? %v", palinDromeCheck2)
	palinDromeCheck3 := IsPalindrome(" ")
	fmt.Printf("\nExplanation: is '' a palindrome? %v", palinDromeCheck3)
	palinDromeCheck4 := IsPalindrome("Madam I'm Adam")
	fmt.Printf("\nExplanation: is 'Madam I'm Adam' a palindrome? %v", palinDromeCheck4)

	fmt.Printf("\nString Fields demo:")
	BasicFieldsDemo()
	MoreComplicatedFieldsDemo()

	romanExample1 := RomanNumeralsToNumber("III")
	fmt.Printf("\nRoman Example 1: III should be 3 [%d]", romanExample1)
	romanExample2 := RomanNumeralsToNumber("LVIII")
	fmt.Printf("\nRoman Example 1: LVIII should be 58 [%d]", romanExample2)
	romanExample3 := RomanNumeralsToNumber("MCMXCIV")
	fmt.Printf("\nRoman Example 1: MCMXCIV should be 1994 [%d]", romanExample3)

	reverseRomanExample1a := ReverseRoman(3749)
	fmt.Printf("\nReverseRoman Example 1s: 3749 should be MMMDCCXLIX [%s]", reverseRomanExample1a)

	fmt.Printf("\nAlternating Parity Check 12345 should be true [%v]", ParityCheckString("12345"))
	fmt.Printf("\nAlternating Parity Check 123457 should be false [%v]", ParityCheckString("123457"))
	fmt.Printf("\nAlternating Parity Check 1324 should be false [%v]", ParityCheckString("1324"))
	fmt.Printf("\nAlternating Parity Check 1032 should be true [%v]", ParityCheckSlice([]int{1, 0, 3, 2}))
	fmt.Printf("\nAlternating Parity Check 0 should be true [%v]", ParityCheckSlice([]int{0}))

	fmt.Printf("\nReverse String '123ABC'  [%v]", basics.ReverseString("123ABC"))
	fmt.Printf("\nString of numbers '1234' to slice of integers [%v]", basics.StringToSlice("1234"))

	fmt.Printf("\nNumber to alphabet representation '1' [%v]", basics.NumberToAlphabet(1))
	fmt.Printf("\nNumber to alphabet representation '23' [%v]", basics.NumberToAlphabet(23))
	fmt.Printf("\nNumber to alphabet representation '26' [%v]", basics.NumberToAlphabet(26))

	fmt.Printf("\nAlpha to number representation 'z' [%v]", basics.AlphabetToNumber('z'))
	fmt.Printf("\nAlpha to number representation 'v' [%v]", basics.AlphabetToNumber('v'))
	fmt.Printf("\nAlpha to number representation 'f' [%v]", basics.AlphabetToNumber('f'))
	fmt.Printf("\nAlpha to number representation 'a' [%v]", basics.AlphabetToNumber('a'))

	fmt.Printf("\nTransform numbers to string '1234567890'  [%v]", basics.TransformNumbersToString("1234567890"))
	fmt.Printf("\nTransform string to numbers 'abcdefghijklm'  [%v]", basics.TransformStringToNumbers("abcdefghijklm"))

	fmt.Printf("\nMinumum Window: Example 1 s = 'ADOBECODEBANC', t = 'ABC' , answer[%v]", MinimumWindow("ADOBECODEBANC", "ABC"))
	fmt.Printf("\nMinumum Window: Example 2 s = 'a', t = 'a' , answer[%v]", MinimumWindow("a", "a"))
	fmt.Printf("\nMinumum Window: Example 3 s = 'a', t = 'aa' , answer[%v]", MinimumWindow("a", "aa"))

	fmt.Printf("\nLuhnCheck '79927398713' should be true, answer [%v]", LuhnCheck("79927398713"))
	fmt.Printf("\nLuhnCheck '49927398716' answer [%v]", LuhnCheck("49927398716"))
	fmt.Printf("\nLuhnCheck '49927398717' answer [%v]", LuhnCheck("49927398717"))
	fmt.Printf("\nLuhnCheck '1234567812345678' answer [%v]", LuhnCheck("1234567812345678"))
	fmt.Printf("\nLuhnCheck '1234567812345670' answer should be true[%v]", LuhnCheck("1234567812345670"))

	fmt.Printf("\nBinary Search should find at index 5 [%v]", BinarySearch([]int{-6, 0, 1, 2, 5, 6, 8, 9}, 6))
	fmt.Printf("\nBinary Search should find at index 0 [%v]", BinarySearch([]int{-6, 0, 1, 2, 5, 6, 8, 9}, -6))
	fmt.Printf("\nBinary Search should find at index 7 [%v]", BinarySearch([]int{-6, 0, 1, 2, 5, 6, 8, 9}, 9))
	fmt.Printf("\nBinary Search should not find [%v]", BinarySearch([]int{-6, 0, 1, 2, 5, 6, 8, 9}, -10))

	fmt.Printf("\nMedian should be 3 [%v]", Median([]int{1, 2, 3, 4, 5}))
	fmt.Printf("\nMedian should be 3 [%v]", Median([]int{5, 2, 1, 4, 3}))
	fmt.Printf("\nMedian should be 1 [%v]", Median([]int{1}))
	fmt.Printf("\nMedian should be 2 [%v]", Median([]int{1, 3}))
	fmt.Printf("\nMedian should be 4 [%v]", Median([]int{1, 2, 3, 4, 5, 6, 7, 8}))

}

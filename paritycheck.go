package main

import (
	"practice2024/basics"
	"strconv"
)

/*  given an array of numbers, check if they have alternating parity

examples
[1,2,3,4] # odd, even, odd, even (True, because they alternate)
[1,3,2,4] # odd, odd, even, even (False, because they don't alternate)

also deal with any conversions
*/

func ParityCheckSlice(slice []int) bool {

	for i := 0; i < len(slice)-1; i++ {

		if isOdd(slice[i]) == isOdd(slice[i+1]) {
			return false
		}

	}

	//if we get here its ok
	return true
}

func isOdd(i int) bool {
	return i%2 != 0
}

// func isEven(i int) bool {
// 	return i%2 == 0
// }

func ParityCheckString(s string) bool {
	return ParityCheckSlice(basics.StringToSlice(s))
}

func ParityCheckNumber(num int) bool {
	return ParityCheckSlice(basics.StringToSlice(strconv.Itoa(num)))
}

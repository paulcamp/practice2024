package main

import "strconv"

//given a number such as 12 , return the sum of the numbers 1+2 . In this case the answer would be 3

func addNumbers(in int) int {

	strNum := strconv.Itoa(in)

	sum := 0
	for _, v := range strNum {
		sum += int(v - '0')

	}
	return sum

}

package main

import (
	"sort"
)

//find the median of a slice of ints

func Median(in []int) int {
	if len(in) == 0 {
		return 0
	}

	//sort the slice
	sort.Slice(in, func(i, j int) bool {
		return in[i] < in[j]
	})

	//find the mid point
	mid := len(in) / 2

	//example 12345
	//length 5/2 = 2.5 this will round down to 2, which is the correct index
	//index: 0 1 2 3 4
	//value: 1 2 3 4 5
	//answer is 3

	//if theres an even number of elements then find mid value between the 2 middle values
	//example 13
	//length 2/2 = 1
	//median = (1+3) /2 = 2

	if len(in)%2 == 0 {
		median := (in[mid-1] + in[mid]) / 2
		return median
	}

	return in[mid]

}

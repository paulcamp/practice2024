package main

import (
	"fmt"
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
	//12345
	//5/2 = 2.5
	fmt.Printf("\nMid is [%d]", mid)

	//1234
	//4/2 = 2  not sure what answer should be...

	if len(in)%2 == 0 {
		median := (in[mid-1] + in[mid]) / 2
		return median
	}

	return in[mid]
	//if theres an even number of elements then find mid value between the 2 middle values

}

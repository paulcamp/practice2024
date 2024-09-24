package main

import (
	"fmt"
	"strings"
)

//practice using string.Fields and FieldsFunc

func BasicFieldsDemo() {
	fmt.Println()
	fmt.Println(basicSplit("Cat  Dog    Bird Wolf"))
	fmt.Println(splitOnAny("Cat,Dog;Bird:Wolf"))
}

func basicSplit(s string) []string {
	return strings.Fields(s)
}

func splitOnAny(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool { return r == ';' || r == ',' || r == ':' })

	//or using standalone    f := func(c rune) bool {
	//    return c == ';' || c == ',' || c == ':'
	//}
}

func MoreComplicatedFieldsDemo() {

	lines := []string{
		"Time:        44     70     70     80",
		"Distance:   283   1134   1134   1491",
	}

	timesRow, distanceRow := getTimesAndDistances(lines)

	fmt.Println(timesRow)
	fmt.Println(distanceRow)

}

func getTimesAndDistances(lines []string) ([]string, []string) {
	timesRow := strings.Fields(lines[0])[1:] //ignore 1st thing which will be "Time:"
	distanceRow := strings.Fields(lines[1])[1:]
	return timesRow, distanceRow
}

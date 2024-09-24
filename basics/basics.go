package basics

import (
	"regexp"
	"strconv"
	"strings"
)

//package intended for helper style funcs

func ReverseString(s string) string {
	//note this wont cope with special characters,
	//look at https://rosettacode.org/wiki/Reverse_a_string#Go for that functionality
	//strings.Map()
	r := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		r[i] = s[len(s)-1-i]
	}
	return string(r)
}

// efficient number to alphabet representation
const abc = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func NumberToAlphabet(i int) string {
	return abc[i-1 : i]
}

var AlphabetMap map[rune]int = map[rune]int{
	'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8, 'i': 9, 'j': 10,
	'k': 11, 'l': 12, 'm': 13, 'n': 14, 'o': 15, 'p': 16, 'q': 17, 'r': 18, 's': 19,
	't': 20, 'u': 21, 'v': 22, 'w': 23, 'x': 24, 'y': 25, 'z': 26,
}

func AlphabetToNumber(c rune) int {

	return AlphabetMap[c]
}

// Transform string 123 to letters ABC
// Note: limited to first 9 numbers '1234567890'  [abcdefghi`]
func TransformNumbersToString(s string) string {

	mapped := func(r rune) rune {
		return r + '0'
	}

	return strings.Map(mapped, s)
}

// Transform string abc to 123
// Note: limited to first 9 numbers 'abcdefghijklm'  [123456789:;<=]
func TransformStringToNumbers(s string) string {

	mapped := func(r rune) rune {
		return r - '0'
	}

	return strings.Map(mapped, s)
}

func StripNonAlphaNumericWithRegEx(in string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9 ]+")
	return reg.ReplaceAllString(in, "")
}

func StripNonAlphaNumericAndSpaces(in string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	return reg.ReplaceAllString(in, "")
}

func StripNonAlphaNumericWithLoop(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') ||
			b == ' ' {
			result.WriteByte(b)
		}
	}
	return result.String()
}

func SliceToString(a []int) string {
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}

	return strings.Join(b, ",")
}

// want "1234567890" >> [1 2 3 4 5 6 7 8 9 0]
func StringToSlice(s string) []int {
	chars := []byte(s)
	var result []int
	for _, v := range chars {
		num, _ := strconv.Atoi(string(v))
		result = append(result, num)

	}

	return result
}

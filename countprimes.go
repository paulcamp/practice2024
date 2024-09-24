package main

import (
	"math"
)

// Given an integer n, return the number of prime numbers that are strictly less than n.
// a prime number is whole number greater than 1 that cannot be exactly divided by any whole number other than itself and 1 (e.g. 2, 3, 5, 7, 11).

func CountPrimesBasic(n int) int {
	count := 0
	for i := 0; i <= n; i++ {
		if isPrime(i) {
			//fmt.Println(i)
			count++
		}
	}
	return count
}

// try dividing by each number below it

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// skip all even numbers apart from 2
// only check odd numbers because we know all even numbers are divisble by 2
// finally only check up to and including the square root of the number, because higher than that would be mathematically impossible
func CountPrimesOptimised(n int) int {
	count := 0
	for i := 0; i <= n; i++ {
		if isPrimeOptimized(i) {
			//fmt.Println(i)
			count++
		}
	}
	return count
}

func isPrimeOptimized(n int) bool {
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	}

	sqrt := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrt; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

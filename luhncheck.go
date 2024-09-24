package main

/*The Luhn algorithm, also known as the modulus 10 or mod 10 algorithm,
is a simple checksum formula used to validate a variety of identification numbers,
 such as credit card numbers, IMEI numbers, Canadian Social Insurance Numbers.
  The LUHN formula was created in the late 1960s by a group of mathematicians.
  Shortly thereafter, credit card companies adopted it.
  Because the algorithm is in the public domain, it can be used by anyone.
  Most credit cards and many government identification numbers use the algorithm
  as a simple method of distinguishing valid numbers from mistyped or otherwise incorrect numbers.
  It was designed to protect against accidental errors, not malicious attacks

Rosetta Code Alogo Desciption:
Reverse the order of the digits in the number.
Take the first, third, ... and every other odd digit in the reversed digits and sum them to form the partial sum s1
Taking the second, fourth ... and every other even digit in the reversed digits:
Multiply each digit by two and sum the digits if the answer is greater than nine to form partial sums for the even digits
Sum the partial sums of the even digits to form s2
If s1 + s2 ends in zero then the original number is in the form of a valid credit card number as verified by the Luhn test.

*/

func LuhnCheck(s string) bool {

	isSecond := false
	sum := 0

	for i := len(s) - 1; i >= 0; i-- {

		digit := int(s[i] - '0') //does this even work in GO - yes it does

		if isSecond {
			digit = digit * 2
			if digit > 9 { //eg. 10 = 1 , 12 = 3, 18 = 9
				digit -= 9
			}
		}

		sum += digit

		isSecond = !isSecond //alternate the flag
	}

	return sum%10 == 0
}

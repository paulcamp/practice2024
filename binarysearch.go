package main

/*Given the starting point of a range, the ending point of a range, and the "secret value",
implement a binary search through a sorted integer array for a certain number.
 Implementations can be recursive or iterative (both if you can).

 Print out whether or not the number was in the array afterwards. If it was, print the index also.

*/

func BinarySearch(in []int, find int) int {

	low := 0
	high := len(in) - 1

	for low <= high {

		mid := (low + high) / 2
		if in[mid] > find {
			high = mid - 1
		} else if in[mid] < find {
			low = mid + 1
		} else {
			return mid
		}
	}

	return -1 //not found
}

/*Additional:
https://rosettacode.org/wiki/Binary_search
The following algorithms return the leftmost place where the given element can be correctly inserted
(and still maintain the sorted order).
This is the lower (inclusive) bound of the range of elements that are equal to the given value (if any).
 Equivalently, this is the lowest index where the element is greater than or equal to the given value (since if it were any lower, it would violate the ordering), or 1 past the last index if such an element does not exist. This algorithm does not determine if the element is actually found. This algorithm only requires one comparison per level.
*/

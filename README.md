# Divide and conquer

1. [+30%] A unimodal array is an array that has a sequence of monotonically increasing integers
   followed by a sequence of monotonically decreasing integers.

   `unimodalMax(in []int) int` is a function that finds the max element in a unimodal array
   implemented as [follows](https://go.dev/play/p/buvERC5yZYG) (discussed during the lesson).

   Write a recurrence equation for the algorithm that `unimodalMax` implements. Solve the equation
   and provide the time complexity for the algorithm expressed in $O(...)$ notation.

2. [+70%] You are given an integer array of length $n$. Write a program that finds the contiguous
   subarray within the array having the largest sum.

   The program must return an integer representing the maximum possible sum of the contiguous subarray.

   The program must implement a divide-and-conquer algorithm (discussed in the lesson):

     1. Divide the given array in two halves.
     2. Return the maximum of following three

        * maximum subarray sum in the left half
        * maximum subarray sum in the right half
        * maximum subarray sum that includes theright-most element in the left half and the left-most
          element in the right half.

   Write a recurrence equation for the algorithm, solve the equation and provide the time complexity for
   the algorithm expressed in $O(...)$ notation.

   **Example 1**

   ```
   Input: [-1 1]
   Output: 1
   ```

   The subarray `[1]` contains the max sum.

   **Exam// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}ple 2**

   ```
   Input: [-1 1 2]
   Output: 3
   ```

   The subarray `[1 2]` contains the max sum.

   **Example 3**

   ```
   Input: [-2 1 -3 4 -1 2 1 -5 4]
   Output: 6
   ```

   The subarray `[4 -1 2 1]` contains the max sum.

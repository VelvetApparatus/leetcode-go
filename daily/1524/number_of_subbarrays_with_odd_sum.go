package _524

/*


	Given an array of integers arr, return the number of subarrays with an odd sum.
	Since the answer can be very large, return it modulo 1e9 + 7.

	Example 1:

	Input: arr = [1,3,5]
	Output: 4
	Explanation: All subarrays are [[1],[1,3],[1,3,5],[3],[3,5],[5]]
	All sub-arrays sum are [1,4,9,3,8,5].
	Odd sums are [1,9,3,5] so the answer is 4.

	URL: https://leetcode.com/problems/number-of-sub-arrays-with-odd-sum/description/
*/

func numOfSubarrays(arr []int) int {
	sum := 0
	odd := 0
	even := 1 // empty list is even number

	for _, a := range arr {
		sum += a
		if sum&1 == 1 {
			odd++
		} else {
			even++
		}
	}
	return (odd * even) % (1e9 + 7)
}

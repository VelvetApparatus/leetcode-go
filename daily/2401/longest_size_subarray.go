package _2401

/*

	You are given an array nums consisting of positive integers.

	We call a subarray of nums nice if the bitwise AND of every pair of elements that are in different positions in the subarray is equal to 0.

	Return the length of the longest nice subarray.

	A subarray is a contiguous part of an array.

	Note that subarrays of length 1 are always considered nice.



	Example 1:

	Input: nums = [1,3,8,48,10]
	Output: 3
	Explanation: The longest nice subarray is [3,8,48]. This subarray satisfies the conditions:
	- 3 AND 8 = 0.
	- 3 AND 48 = 0.
	- 8 AND 48 = 0.
	It can be proven that no longer nice subarray can be obtained, so we return 3.


	URL: https://leetcode.com/problems/longest-nice-subarray/description/?envType=daily-question&envId=2025-03-18

*/

func longestNiceSubarray(nums []int) int {
	var (
		best      = 1
		bitwiseOr = 0
		left      = 0
		right     = len(nums)
	)

	for right = 0; right < len(nums); right++ {
		for (bitwiseOr & nums[right]) != 0 {
			bitwiseOr ^= nums[left]
			left++
		}

		bitwiseOr |= nums[right]

		best = max(best, right-left+1)
	}

	return best
}

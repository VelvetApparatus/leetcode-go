package _749

/*
	You are given an integer array nums.
	The absolute sum of a subarray [numsl, numsl+1, ..., numsr-1, numsr] is abs(numsl + numsl+1 + ... + numsr-1 + numsr).
	Return the maximum absolute sum of any (possibly empty) subarray of nums.
	Note that abs(x) is defined as follows:
	If x is a negative integer, then abs(x) = -x.
	If x is a non-negative integer, then abs(x) = x.

	Example 1:

	Input: nums = [1,-3,2,3,-4]
	Output: 5
	Explanation: The subarray [2,3] has absolute sum = abs(2+3) = abs(5) = 5.

	URL: https://leetcode.com/problems/maximum-absolute-sum-of-any-subarray/description/
*/

func maxAbsoluteSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var (
		maxEnding = 0
		minEnding = 0
		res       = 0
	)

	for _, num := range nums {
		maxEnding = max(num, maxEnding+num)
		minEnding = min(num, minEnding+num)

		res = max(res, max(maxEnding, -minEnding))
	}
	return res
}

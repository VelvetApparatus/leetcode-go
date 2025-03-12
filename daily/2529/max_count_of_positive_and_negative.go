package _2529

/*

	Given an array nums sorted in non-decreasing order, return the maximum between the number of positive integers and the number of negative integers.

	In other words, if the number of positive integers in nums is pos and the number of negative integers is neg, then return the maximum of pos and neg.
	Note that 0 is neither positive nor negative.

	Example 1:

	Input: nums = [-2,-1,-1,1,2,3]
	Output: 3
	Explanation: There are 3 positive integers and 3 negative integers. The maximum count among them is 3.

	URL: https://leetcode.com/problems/maximum-count-of-positive-integer-and-negative-integer/description/?envType=daily-question&envId=2025-03-12

*/

func maximumCount(nums []int) int {
	pos, neg := 0, 0

	for i := range nums {
		if nums[i] > 0 {
			pos++
			continue
		}
		if nums[i] < 0 {
			neg++
			continue
		}
	}
	return max(pos, neg)
}

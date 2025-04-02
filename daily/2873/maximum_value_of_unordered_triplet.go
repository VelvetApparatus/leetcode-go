package _2873

/*

	You are given a 0-indexed integer array nums.

	Return the maximum value over all triplets of indices (i, j, k) such that i < j < k. If all such triplets have a negative value, return 0.

	The value of a triplet of indices (i, j, k) is equal to (nums[i] - nums[j]) * nums[k].



	Example 1:

	Input: nums = [12,6,1,2,7]
	Output: 77
	Explanation: The value of the triplet (0, 2, 4) is (nums[0] - nums[2]) * nums[4] = 77.
	It can be shown that there are no ordered triplets of indices with a value greater than 77.

	URL: https://leetcode.com/problems/maximum-value-of-an-ordered-triplet-i/description/?envType=daily-question&envId=2025-04-02

*/

func maximumTripletValue(nums []int) int64 {
	n := len(nums)
	if n < 3 {
		return 0
	}

	res := int64(0)

	leftMax := make([]int, n)
	leftMax[0] = nums[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], nums[i])
	}

	rightMax := make([]int, n)
	rightMax[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], nums[i])
	}

	for j := 1; j < n-1; j++ {
		iMax := leftMax[j-1]
		kMax := rightMax[j+1]

		res = max(res, int64((iMax-nums[j])*kMax))
	}

	return res
}

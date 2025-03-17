package _2206

/*

	You are given an integer array nums consisting of 2 * n integers.

	You need to divide nums into n pairs such that:

	Each element belongs to exactly one pair.
	The elements present in a pair are equal.
	Return true if nums can be divided into n pairs, otherwise return false.



	Example 1:

	Input: nums = [3,2,3,2,2,2]
	Output: true
	Explanation:
	There are 6 elements in nums, so they should be divided into 6 / 2 = 3 pairs.
	If nums is divided into the pairs (2, 2), (3, 3), and (2, 2), it will satisfy all the conditions.

	URL: https://leetcode.com/problems/divide-array-into-equal-pairs/description/?envType=daily-question&envId=2025-03-17

*/

func divideArray(nums []int) bool {
	m := make(map[int]bool, len(nums))
	pairs := 0
	for _, v := range nums {
		if m[v] {
			m[v] = false
			pairs++
		} else {
			m[v] = true
		}
	}
	return pairs == len(nums)/2
}

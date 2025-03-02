package _570

import "slices"

/*

	You are given two 2D integer arrays nums1 and nums2.

	nums1[i] = [idi, vali] indicate that the number with the id idi has a value equal to vali.
	nums2[i] = [idi, vali] indicate that the number with the id idi has a value equal to vali.
	Each array contains unique ids and is sorted in ascending order by id.

	Merge the two arrays into one array that is sorted in ascending order by id, respecting the following conditions:

	Only ids that appear in at least one of the two arrays should be included in the resulting array.
	Each id should be included only once and its value should be the sum of the values of this id in the two arrays.
	If the id does not exist in one of the two arrays, then assume its value in that array to be 0.
	Return the resulting array. The returned array must be sorted in ascending order by id.


	Example 1:

	Input: nums1 = [[1,2],[2,3],[4,5]], nums2 = [[1,4],[3,2],[4,1]]
	Output: [[1,6],[2,3],[3,2],[4,6]]
	Explanation: The resulting array contains the following:
	- id = 1, the value of this id is 2 + 4 = 6.
	- id = 2, the value of this id is 3.
	- id = 3, the value of this id is 2.
	- id = 4, the value of this id is 5 + 1 = 6.

	URL: https://leetcode.com/problems/merge-two-2d-arrays-by-summing-values/description/

*/

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	m := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		m[nums1[i][0]] += nums1[i][1]
	}
	for i := 0; i < len(nums2); i++ {
		m[nums2[i][0]] += nums2[i][1]
	}
	res := make([][]int, 0)
	for k, v := range m {
		res = append(res, []int{k, v})
	}

	slices.SortFunc(res, func(a, b []int) int {
		if a[0] < b[0] {
			return -1
		}
		return 1
	})
	return res
}

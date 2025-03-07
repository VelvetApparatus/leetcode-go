package _2965

/*

	You are given a 0-indexed 2D integer matrix grid of size n * n with values in the range [1, n2].
	Each integer appears exactly once except a which appears twice and b which is missing.
	The task is to find the repeating and missing numbers a and b.

	Return a 0-indexed integer array ans of size 2 where ans[0] equals to a and ans[1] equals to b.

	Example 1:

	Input: grid = [[1,3],[2,2]]
	Output: [2,4]
	Explanation: Number 2 is repeated and number 4 is missing so the answer is [2,4].

	URL: https://leetcode.com/problems/find-missing-and-repeated-values/description/?envType=daily-question&envId=2025-03-06

*/

func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	size := n * n
	count := make([]int, size+1)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			count[grid[i][j]]++
		}
	}

	a, b := -1, -1

	for num := 1; num <= size; num++ {
		if count[num] == 2 {
			a = num
		} else if count[num] == 0 {
			b = num
		}
	}

	return []int{a, b}
}

package _718

/*
	Given an integer n, find a sequence with elements in the range [1, n] that satisfies all of the following:

	The integer 1 occurs once in the sequence.
	Each integer between 2 and n occurs twice in the sequence.
	For every integer i between 2 and n, the distance between the two occurrences of i is exactly i.
	The distance between two numbers on the sequence, a[i] and a[j], is the absolute difference of their indices, |j - i|.

	Return the lexicographically largest sequence. It is guaranteed that under the given constraints, there is always a solution.

	A sequence a is lexicographically larger than a sequence b (of the same length) if in the first position where a and b differ,
	sequence a has a number greater than the corresponding number in b. For example, [0,1,9,0] is lexicographically larger than [0,1,5,6]
	because the first position they differ is at the third number, and 9 is greater than 5.

	Example 1:

	Input: n = 3
	Output: [3,1,2,3,2]
	Explanation: [2,3,2,1,3] is also a valid sequence, but [3,1,2,3,2] is the lexicographically largest valid sequence.

	URL: https://leetcode.com/problems/construct-the-lexicographically-largest-valid-sequence/description/?envType=daily-question&envId=2025-02-16
*/

func constructDistancedSequence(n int) []int {
	res := make([]int, 2*n-1)
	// bitmap for tracking used integers
	used := make([]bool, n+1)
	backtrack(res, used, n, 0)
	return res
}

func backtrack(res []int, used []bool, n, idx int) bool {
	// skip filled cells
	for idx < len(res) && res[idx] != 0 {
		idx++
	}

	// there is a decision
	if idx == len(res) {
		return true
	}

	// We try to put numbers from n to 1 (to get the lexicographically largest sequence)
	for num := n; num >= 1; num-- {
		if used[num] {
			continue
		}

		if num == 1 {
			res[idx] = 1
			used[1] = true
			if backtrack(res, used, n, idx+1) {
				return true
			}
			res[idx] = 0
			used[1] = false
		} else {
			if idx+num < len(res) && res[idx] == 0 && res[idx+num] == 0 {
				res[idx], res[idx+num] = num, num
				used[num] = true
				if backtrack(res, used, n, idx+1) {
					return true
				}
				res[idx], res[idx+num] = 0, 0
				used[num] = false
			}
		}
	}

	return false
}

package _79

/*

	A sequence x1, x2, ..., xn is Fibonacci-like if:
	n >= 3
	xi + xi+1 == xi+2 for all i + 2 <= n
	Given a strictly increasing array arr of positive integers forming a sequence,
	return the length of the longest Fibonacci-like subsequence of arr. If one does not exist, return 0.

	A subsequence is derived from another sequence arr by deleting any number of elements (including none) from arr,
	without changing the order of the remaining elements. For example, [3, 5, 8] is a subsequence of [3, 4, 5, 6, 7, 8].

	Example 1:

	Input: arr = [1,2,3,4,5,6,7,8]
	Output: 5
	Explanation: The longest subsequence that is fibonacci-like: [1,2,3,5,8].

	URL: https://leetcode.com/problems/length-of-longest-fibonacci-subsequence/description/

*/

// IntPair represents a pair of numbers (y, diff) used in the DP map.
type IntPair struct {
	first  int
	second int
}

func lenLongestFibSubseq(arr []int) int {
	maxLen := 0
	dp := make(map[IntPair]int)    // DP map to store Fibonacci-like subsequence lengths
	seen := make(map[int]struct{}) // HashSet to track seen elements

	for k, z := range arr {
		seen[z] = struct{}{} // Mark z as seen

		// Iterate over all previous values y in arr[..k] in reverse order
		for j := k - 1; j >= 0; j-- {
			y := arr[j]
			diff := z - y

			// Early break: If diff is too large, further iterations are invalid
			if diff >= y {
				break
			}

			// Check if diff exists in the set
			if _, found := seen[diff]; found {
				// Get previous sequence length, defaulting to 2 if not found
				curLen := max(dp[IntPair{y, diff}], 2) + 1
				maxLen = max(maxLen, curLen)

				// Store computed length for (z, y)
				dp[IntPair{z, y}] = curLen
			}
		}
	}

	return maxLen
}

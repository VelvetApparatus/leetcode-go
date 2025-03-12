package _1358

/*

	Given a string s consisting only of characters a, b and c.

	Return the number of substrings containing at least one occurrence of all these characters a, b and c.


	Example 1:

	Input: s = "abcabc"
	Output: 10
	Explanation: The substrings containing at least one occurrence
	of the characters a, b and c are "abc", "abca", "abcab", "abcabc", "bca", "bcab", "bcabc", "cab", "cabc" and "abc" (again).

	URL: https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/description/?envType=daily-question&envId=2025-03-11

*/

func numberOfSubstrings(s string) int {
	freq := map[byte]int{'a': 0, 'b': 0, 'c': 0}
	left, ans := 0, 0

	for right := 0; right < len(s); right++ {
		freq[s[right]]++

		for freq['a'] > 0 && freq['b'] > 0 && freq['c'] > 0 {
			ans += len(s) - right
			freq[s[left]]--
			left++
		}
	}

	return ans
}

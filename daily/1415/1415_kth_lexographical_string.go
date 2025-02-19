package _415

/*
	A happy string is a string that:

	consists only of letters of the set ['a', 'b', 'c'].
	s[i] != s[i + 1] for all values of i from 1 to s.length - 1 (string is 1-indexed).
	For example, strings "abc", "ac", "b" and "abcbabcbcb" are all happy strings and strings "aa", "baa" and "ababbc" are not happy strings.

	Given two integers n and k, consider a list of all happy strings of length n sorted in lexicographical order.

	Return the kth string of this list or return an empty string if there are less than k happy strings of length n.

	Example 3:

	Input: n = 3, k = 9
	Output: "cab"
	Explanation: There are 12 different happy string of length 3
	["aba", "abc", "aca", "acb", "bab", "bac", "bca", "bcb", "cab", "cac", "cba", "cbc"].
	You will find the 9th string = "cab"

	URL: https://leetcode.com/problems/the-k-th-lexicographical-string-of-all-happy-strings-of-length-n/description/
*/

func getHappyString(n int, k int) string {
	set := []uint8{'a', 'b', 'c'}
	res := make([]string, 0, n)
	kk := k

	for i := range set {
		kk = backtrace(string(set[i]), n-1, set, &res, kk)
	}

	if kk > 0 {
		return ""
	}

	return res[k-1]
}

func backtrace(
	s string,
	n int,
	set []uint8,
	res *[]string,
	k int,
) int {

	if k == 0 {
		return 0
	}

	if n == 0 {
		*res = append(*res, s)
		return k - 1
	}

	for i := 0; i < len(set); i++ {
		if s[len(s)-1] == set[i] {
			continue
		}
		k = backtrace(s+string(set[i]), n-1, set, res, k)
	}
	return k
}

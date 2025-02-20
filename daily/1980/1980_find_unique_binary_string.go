package _980

/*

	Given an array of strings nums containing n unique binary strings each of length n,
	return a binary string of length n that does not appear in nums.
	If there are multiple answers, you may return any of them.

	Example 1:

	Input: nums = ["01","10"]
	Output: "11"
	Explanation: "11" does not appear in nums. "00" would also be correct.

	URL: https://leetcode.com/problems/find-unique-binary-string/description/

*/

func findDifferentBinaryString(nums []string) string {
	set := map[string]bool{}
	for i := 0; i < len(nums); i++ {
		set[nums[i]] = true
	}
	str, found := backtrace("", len(nums[0]), set)
	if found {
		return str
	}
	return ""
}

func backtrace(
	s string,
	n int,
	set map[string]bool,
) (string, bool) {
	if n == 0 {
		_, found := set[s]
		return s, !found
	}

	str, found := backtrace(s+"0", n-1, set)
	if found {
		return str, true
	}

	str, found = backtrace(s+"1", n-1, set)
	if found {
		return str, true
	}
	return "", false
}

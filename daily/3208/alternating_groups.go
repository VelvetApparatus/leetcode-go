package _3208

/*

	There is a circle of red and blue tiles. You are given an array of integers colors and an integer k.
	The color of tile i is represented by colors[i]:

	colors[i] == 0 means that tile i is red.
	colors[i] == 1 means that tile i is blue.
	An alternating group is every k contiguous tiles in the circle with alternating colors
	(each tile in the group except the first and last one has a different color from its left and right tiles).

	Return the number of alternating groups.

	Note that since colors represents a circle, the first and the last tiles are considered to be equal to each other.


	Example 1:

	Input: colors = [0,1,0,1,0], k = 3

	Output: 3

	URL: https://leetcode.com/problems/alternating-groups-ii/description/?envType=daily-question&envId=2025-03-09

*/

func numberOfAlternatingGroups(colors []int, k int) int {
	l, r := 0, 0
	res := 0
	n := len(colors)
	for l < n {
		r++
		if colors[r%n] == colors[(r-1)%n] {
			l = r
		} else if r-l+1 == k {
			res++
			l++
		}
	}
	return res
}

package _2579

/*

	There exists an infinitely large two-dimensional grid of uncolored unit cells. You are given a positive integer n, indicating that you must do the following routine for n minutes:

	At the first minute, color any arbitrary unit cell blue.
	Every minute thereafter, color blue every uncolored cell that touches a blue cell.
	Below is a pictorial representation of the state of the grid after minutes 1, 2, and 3.

	Return the number of colored cells at the end of n minutes.

	Example 1:

	Input: n = 1
	Output: 1
	Explanation: After 1 minute, there is only 1 blue cell, so we return 1.


	URL: https://leetcode.com/problems/count-total-number-of-colored-cells/description/

*/

func coloredCells(n int) int64 {
	res := int64(0)
	d := int64(n*2 - 1)
	res += d
	d -= 2
	for i := 1; i < n; i++ {
		res += d * 2
		d -= 2
	}
	return res
}

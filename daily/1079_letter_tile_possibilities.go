package daily

/*
	You have n  tiles, where each tile has one letter tiles[i] printed on it.
	Return the number of possible non-empty sequences of letters you can make using the letters printed on those tiles.


	Example 1:

	Input: tiles = "AAB"
	Output: 8
	Explanation: The possible sequences are "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA".


	URL: https://leetcode.com/problems/letter-tile-possibilities/?envType=daily-question&envId=2025-02-17
*/

func numTilePossibilities(tiles string) int {
	count := make(map[rune]int)
	for _, tile := range tiles {
		count[tile]++
	}
	return backtrack(count)
}

func backtrack(count map[rune]int) int {
	total := 0
	for ch, v := range count {
		if v > 0 {
			count[ch]--
			total += 1 + backtrack(count)
			count[ch]++
		}
	}
	return total
}

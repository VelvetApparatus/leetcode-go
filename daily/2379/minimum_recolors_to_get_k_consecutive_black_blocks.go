package _2379

import "strings"

func minimumRecolors(blocks string, k int) int {
	res := len(blocks)
	for i := 0; i+k <= len(blocks); i++ {
		res = min(res, strings.Count(blocks[i:i+k], "W"))
	}
	return res
}

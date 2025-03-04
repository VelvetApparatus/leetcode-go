package _780

func checkPowersOfThree(n int) bool {
	pow := 1
	for n >= pow {
		pow *= 3
	}
	pow /= 3

	for pow > 0 {
		if n >= pow {
			n -= pow
		}
		pow /= 3
	}
	return n == 0
}

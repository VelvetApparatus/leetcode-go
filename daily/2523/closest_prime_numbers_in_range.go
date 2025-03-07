package _2523

/*

	Given two positive integers left and right, find the two integers num1 and num2 such that:

	left <= num1 < num2 <= right .
	Both num1 and num2 are prime numbers.
	num2 - num1 is the minimum amongst all other pairs satisfying the above conditions.
	Return the positive integer array ans = [num1, num2].
	If there are multiple pairs satisfying these conditions,
	return the one with the smallest num1 value. If no such numbers exist, return [-1, -1].


	Example 1:
	Input: left = 10, right = 19
	Output: [11,13]
	Explanation: The prime numbers between 10 and 19 are 11, 13, 17, and 19.
	The closest gap between any pair is 2, which can be achieved by [11,13] or [17,19].
	Since 11 is smaller than 17, we return the first pair.

	URL: https://leetcode.com/problems/closest-prime-numbers-in-range/description/?envType=daily-question&envId=2025-03-07

*/

func closestPrimes(left int, right int) []int {
	primes := []int{}
	res := []int{-1, -1}

	for i := left; i <= right; i++ {
		if isPrime(i) {
			primes = append(primes, i)
			if len(primes) > 1 {
				if res[0] == -1 {
					res[0] = primes[len(primes)-2]
					res[1] = primes[len(primes)-1]
					continue
				}
				if res[1]-res[0] > primes[len(primes)-1]-primes[len(primes)-2] {
					res[0] = primes[len(primes)-2]
					res[1] = primes[len(primes)-1]
				}
			}
		}
	}
	return res
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	d := 2
	for d*d < n {
		if n%d == 0 {
			return false
		}
		d += 1
	}
	if d*d == n {
		return false
	}
	return true
}

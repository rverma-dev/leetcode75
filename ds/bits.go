package ds

func BitMultiply(m, n int) int {
	var res int
	for n > 0 {
		if n&1 == 1 {
			res = res + m
		}
		m = m << 1
		n = n >> 1
	}
	return res
}

func SieveOfEratosthenes(n int) []int {
	var primes []int
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	return primes
}

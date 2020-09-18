package solve

func countPrimes(n int) int {
	if n == 1 || n == 0 {
		return 0
	}



	notPrime := make([]bool, n)
	for i := 2; i <= n / 2; i++ {
		if !notPrime[i] {
			for j := i; i*j < n; j++ {
				notPrime[i*j] = true
			}
		}
	}

	count := 0
	for i := 2; i < n; i++ {
		if !notPrime[i] {
			count++
		}
	}

	return count
}

package p50

func myPow(x float64, n int) float64 {
	if n < 0 {
		return doMyPow(1/x, n)
	} else {
		return doMyPow(x, n)
	}
}

func doMyPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n%2 == 0 {
		return doMyPow(x*x, n/2)
	} else {
		return x * doMyPow(x*x, n/2)
	}
}

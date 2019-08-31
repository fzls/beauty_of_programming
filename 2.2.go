package beauty_of_programming

func GetZeros(n int) int {
	resWithoutZeros := 1
	nZeros := 0

	for i := 1; i <= n; i++ {
		resWithoutZeros *= i
		for resWithoutZeros%10 == 0 {
			resWithoutZeros /= 10
			nZeros++
		}
	}

	return nZeros
}

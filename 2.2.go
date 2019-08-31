package beauty_of_programming

type Result struct {
	ResWithoutZeros int
	NZeros          int
}

var cacheMap = make(map[int]*Result)

func GetZeros(n int) int {
	if n == 0 {
		return 1
	}

	_, nZeros := getZeros(n)

	return nZeros
}

func getZeros(n int) (resWithoutZeros int, nZeros int) {
	if n == 1 {
		return 1, 0
	}

	if cache, ok := cacheMap[n]; ok {
		return cache.ResWithoutZeros, cache.NZeros
	}

	resWithoutZeros, nZeros = getZeros(n - 1)
	res := resWithoutZeros * n
	for res%10 == 0 {
		res /= 10
		nZeros++
	}
	resWithoutZeros = res
	cacheMap[n] = &Result{
		ResWithoutZeros: resWithoutZeros,
		NZeros:          nZeros,
	}

	return
}

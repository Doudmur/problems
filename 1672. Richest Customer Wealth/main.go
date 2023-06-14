package test

func MaximumWealth(accounts [][]int) int {
	res := 0
	for _, arr := range accounts {
		sum := 0
		for _, el := range arr {
			sum += el
		}
		if sum > res {
			res = sum
		}
	}

	return res
}

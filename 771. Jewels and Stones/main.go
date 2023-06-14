package main

func NumJewelsInStones(jewels string, stones string) int {
	sum := 0
	m := make(map[rune]struct{})
	for _, el := range jewels {
		m[el] = struct{}{}
	}
	for _, stone := range stones {
		if _, ok := m[stone]; ok {
			sum++
		}
	}
	return sum
}

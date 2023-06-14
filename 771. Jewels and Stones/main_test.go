package main

import "testing"

func TestNumJewelsInStones(t *testing.T) {
	vals := []struct {
		jewels   string
		stones   string
		expected int
	}{
		{
			jewels:   "zZ",
			stones:   "ZZzZFASFzZZSDZZzZAAS",
			expected: 11,
		},
		{
			jewels:   "aA",
			stones:   "aAAbbbb",
			expected: 3,
		},
		{
			jewels:   "z",
			stones:   "ZZ",
			expected: 0,
		},
	}

	for _, Case := range vals {
		res := NumJewelsInStones(Case.jewels, Case.stones)

		if Case.expected != res {
			t.Errorf("Wrong result %d expect %d", res, Case.expected)
		}
	}
}

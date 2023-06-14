package test

import "testing"

func TestMaximumWealth(t *testing.T) {
	vals := []struct {
		accounts [][]int
		expected int
	}{
		{
			accounts: [][]int{{1, 2, 3}, {3, 2, 1}},
			expected: 6,
		},
		{
			accounts: [][]int{{1, 2, 3}, {3, 2, 1}, {9, 8, 7}},
			expected: 24,
		},
		{
			accounts: [][]int{{1, 0, 0}},
			expected: 1,
		},
		{
			accounts: [][]int{{0, 0}},
			expected: 0,
		},
	}

	for _, Case := range vals {
		res := MaximumWealth(Case.accounts)

		if Case.expected != res {
			t.Errorf("Wrong result %d expect %d", res, Case.expected)
		}
	}
}

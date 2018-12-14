package _2_metrixs

import (
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		num     int
		metrixs [][]int
	}{
		{2, [][]int{{1, 2}, {3, 4}}},
		{3, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
		{4, [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}},
	}

	getResult := func(do func(int) [][]int) {
		for _, testcase := range testcases {
			t.Logf("%#v\n", do(testcase.num))
		}
	}

	getResult(matrix)
}

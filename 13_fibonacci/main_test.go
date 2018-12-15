package _3_fibonacci

import "testing"

func Test(t *testing.T) {
	testcases := []struct {
		num  int
		fnum int
	}{
		{2, 1},
		{3, 2},
		{6, 8},
	}

	getResult := func(do func(int) int) {
		for _, testcase := range testcases {
			if testcase.fnum != do(testcase.num) {
				t.Error("except", testcase.fnum, "get", do(testcase.num))
			}
		}
	}

	getResult(fibonacciIterative)
	getResult(fibonacciRecusive)
	getResult(fibonacciRecusiveMem)
}

package _9_steps

import "testing"

func Test(t *testing.T) {
	testcases := []struct {
		num int
	}{
		{2},
		{3},
		{5},
	}

	getResult := func(do func(int)) {
		for _, testcase := range testcases {
			do(testcase.num)
		}
	}

	getResult(steps)

	getResult(stepsRecusive)
}

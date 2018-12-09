package main

import "testing"

func Test(t *testing.T) {
	testcases := []struct {
		num    int
		result int
	}{
		{10032, 23001},
		{-11003223001, -10032230011},
		{-90, -9},
	}

	getResult := func(do func(num int) int) {
		for _, testcase := range testcases {
			if testcase.result != do(testcase.num) {
				t.Error("except", testcase.result, "get", do(testcase.num))
			}
		}
	}

	getResult(P)
	getResult(P2)
}

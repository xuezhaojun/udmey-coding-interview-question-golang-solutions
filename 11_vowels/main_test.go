package _1_vowels

import "testing"

func Test(t *testing.T) {
	testcases := []struct {
		str string
		num int
	}{
		{"Hi There!", 3},
		{"Why do you ask?", 4},
		{"Why?!", 0},
	}

	getResult := func(do func(string) int) {
		for _, testcase := range testcases {
			if testcase.num != do(testcase.str) {
				t.Errorf("except %d get %d", testcase.num, do(testcase.str))
			}
		}
	}

	getResult(vowels)

	getResult(vowelsReg)
}

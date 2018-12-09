package main

import "testing"

func Test(t *testing.T) {
	testcases := []struct {
		str1   string
		result bool
	}{
		{"hello", false},
		{"abba", true},
		{"世界界世", true},
	}

	getResult := func(do func(str1 string) bool) {
		for _, testcase := range testcases {
			if testcase.result != do(testcase.str1) {
				t.Error("except", testcase.result, "get", do(testcase.str1))
			}
		}
	}

	getResult(check)
}

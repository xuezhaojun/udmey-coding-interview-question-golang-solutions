package main

import "testing"

func Test(t *testing.T) {
	testcases := []struct {
		str    string
		result string
	}{
		{"Hello There?", "e"},
		{"11122334", "1"},
	}

	getResult := func(do func(str string) string) {
		for _, testcase := range testcases {
			if testcase.result != do(testcase.str) {
				t.Error("except", testcase.result, "get", do(testcase.str))
			}
		}
	}

	getResult(maxChars)
}

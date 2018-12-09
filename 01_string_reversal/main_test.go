package main

import "testing"

func Test(t *testing.T) {
	testcases := []struct {
		str1 string
		str2 string
	}{
		{"cda54321", "12345adc"},
		{"hello 世界", "界世 olleh"},
	}

	getResult := func(do func(str1 string) string) {
		for _, testcase := range testcases {
			if testcase.str2 != do(testcase.str1) {
				t.Error("except", testcase.str2, "get", do(testcase.str1))
			}
		}
	}

	getResult(reverse02)

	getResult(reverseByByte)
}

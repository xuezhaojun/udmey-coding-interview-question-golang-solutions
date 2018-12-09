package main

import "testing"

func Test(t *testing.T) {
	testcases := []struct {
		str1   string
		str2   string
		result bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"rail safety", "fairy tales", true},
	}

	getResult := func(isanagram func(str1, str2 string) bool) {
		for _, testcase := range testcases {
			if testcase.result != isanagram(testcase.str1, testcase.str2) {
				t.Error(testcase.str1, testcase.str2, "except", testcase.result, "get", !testcase.result)
			}
		}
	}

	getResult(anagram)
}

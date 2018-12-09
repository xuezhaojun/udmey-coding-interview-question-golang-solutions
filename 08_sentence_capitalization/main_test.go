package _8_sentence_capitalization

import "testing"

func Test(t *testing.T) {
	testcases := []struct {
		str1 string
		str2 string
	}{
		{"anagram", "Anagram"},
		{"rat,hello", "Rat,Hello"},
		{"rail safety", "Rail Safety"},
	}

	getResult := func(do func(string) string) {
		for _, testcase := range testcases {
			if testcase.str2 != do(testcase.str1) {
				t.Error(testcase.str1, testcase.str2, "except", testcase.str2, "get", do(testcase.str1))
			}
		}
	}

	getResult(c)
}

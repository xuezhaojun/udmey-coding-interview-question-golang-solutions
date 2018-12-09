package _8_sentence_capitalization

import "strings"

// 将一个句子中的每一个词的首字母大写

func c(str string) string {
	ss := strings.Split(str, "")
	for index, _ := range ss {
		if index == 0 || ss[index-1] == " " || ss[index-1] == "," {
			if ss[index][0] <= 'z' {
				ss[index] = strings.ToUpper(ss[index])
			}
		}
	}
	return strings.Join(ss, "")
}

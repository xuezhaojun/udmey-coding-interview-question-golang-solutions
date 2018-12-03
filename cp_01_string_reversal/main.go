package cp_01_string_reversal

import "fmt"

func main() {
	oldstr := "12345"
	fmt.Println(reverse01(oldstr))
}

// 遍历字符串，每次获取到的下一个字符，都放在新字符串的第一位
// %c - 单字符
func reverse01(str string) string {
	var newstr string
	for i, _ := range str {
		newstr = fmt.Sprintf("%c%s", str[i], newstr)
	}
	return newstr
}

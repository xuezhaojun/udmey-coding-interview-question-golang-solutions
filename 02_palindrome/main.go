package main

import "fmt"

func main() {
	str := "hello"
	str1 := "abba"
	str2 := "世界界世"
	fmt.Println(check(str))
	fmt.Println(check(str1))
	fmt.Println(check(str2))
}

// check 用于检验字符串是否满足回文的要求
// 可以借鉴对应
func check(str string) bool {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

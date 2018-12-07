package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(P(100))
	fmt.Println(P(101))
	fmt.Println(P(10032))
	fmt.Println(P(-11003223001))
	fmt.Println(P(-90))
}

// P 是否回文
// 转化为字符串后，通过字符串的回文判断方法
func P(num int) int {
	str := strconv.Itoa(num)
	runes := []rune(str)

	// 去掉正负号
	pre := runes[0]
	if pre == '-' {
		runes = runes[1:]
	}

	// 转化回文
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	newNum, _ := strconv.Atoi(string(runes))

	// 添加正负号
	if pre == '-' {
		newNum = -newNum
	}
	return newNum
}

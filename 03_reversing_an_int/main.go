package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(P(10032))
	fmt.Println(P(-11003223001))
	fmt.Println(P(-90))

	fmt.Println(P2(10032))
	fmt.Println(P2(-11003223001))
	fmt.Println(P2(-90))
}

// 转化为字符串后，通过字符串的回文判断方法
func P(num int) int {
	str := strconv.Itoa(num)
	runes := []rune(str)

	// 去掉负号
	if num < 0 {
		runes = runes[1:]
	}

	// 转化回文
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	newNum, _ := strconv.Atoi(string(runes))

	// 添加负号
	if num < 0 {
		newNum = -newNum
	}
	return newNum
}

// 直接对不同的位数进行判断
func P2(num int) (newNum int) {
	for num != 0 {
		temp := newNum*10 + num%10
		if temp/10 != newNum {
			return 0 // 溢出
		}
		newNum, num = temp, num/10
	}
	return
}

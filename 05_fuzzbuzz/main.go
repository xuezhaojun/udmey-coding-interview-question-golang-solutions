package main

import "fmt"

// fuzzbuzz问题：
// 输入一个整数
// 找出其中的fuzz数，3的倍数
// 找出其中的buzz数,5的倍数

// 号称是面试中，经常出现的最简单题
// 对应leetcode 412

func main() {
	fuzzbuzz(15)
}

func fuzzbuzz(num int) {
	for i:=1; i <= num ; i ++ {
		switch {
		case i%3 == 0 && i%5 == 0 :
			fmt.Println("fizzbuzz")
		case i%3 == 0:
			fmt.Println("fizz")
		case i%5 == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(i)
		}
	}
}

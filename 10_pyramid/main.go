package _0_pyramid

import "fmt"

// 和上一个阶梯类似，输出金字塔型的#
// 例如输出3
// "  #  "
// " ### "
// "#####"

// 首当其冲就是递归
// 用递归，每一层和每一层之间的关系非常清晰
// 比迭代要简单的多
func pyamid(n int) {
	// 最底层的#数
	printPy(n*2-1, 0)
}

// i = # 数量
// j = 每边的<space>数量
func printPy(i, j int) {
	if i <= 0 {
		return
	}
	printPy(i-2, j+1) // 每天上一层，#数量减少2，空格数量每边加1
	for index := 0; index < j; index++ {
		fmt.Print(" ")
	}
	for index := 0; index < i; index++ {
		fmt.Print("#")
	}
	for index := 0; index < j; index++ {
		fmt.Print(" ")
	}
	fmt.Println("")
}

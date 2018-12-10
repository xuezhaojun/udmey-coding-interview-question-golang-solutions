package _9_steps

import "fmt"

// 输入一个整数n
// 输出用 # 和 <space> 表示的阶梯
// 比如 输入3
// '#  '
// "## "
// "###"
func steps(n int) {
	// 矩阵相关的先考虑一下横列之间的关系
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j <= i { // 当列大于行的时候，就打应#
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

// 递归的写
// 这可能时史上最简单的递归的应用了
func stepsRecusive(n int) {
	printNumbers(n, 0)
}

func printNumbers(i, j int) {
	if i == 0 {
		return
	}
	// 输出少一个# ，多个<space>
	printNumbers(i-1, j+1)
	for index := 0; index < i; index++ {
		fmt.Print("#")
	}
	for index := 0; index < j; index++ {
		fmt.Print(" ")
	}
	fmt.Println("")
}

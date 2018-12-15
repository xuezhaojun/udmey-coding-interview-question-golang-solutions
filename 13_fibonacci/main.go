package _3_fibonacci

// fibonacci 数列对于理解时间复杂度，非常的典型，非常重要的一道题
// 什么是 斐波那契数列？ 0,1,1,2,3,5,8 ... 除了第一个和第二个数以外，剩下的数都满足 f(n) = f(n-2)+f(n-1)
// 以下有三种解法需要掌握
// 遍历，递归，优化的递归，矩阵
// 输入 num 代表数列中的第num个数, 第0个数是0，第一个数是1，第二个数是1 ...
// 输出 fnum

// 递归的方式 - 时间复杂度 O(2^n)
func fibonacciRecusive(num int) int {
	if num == 1 || num == 2 {
		return 1
	}
	return fibonacciRecusive(num-1) + fibonacciRecusive(num-2)
}

// 通过缓存已经计算好的结果，然后直接返回，时间复杂度可以优化到 O(n)
var resultMap = make(map[int]int)

func fibonacciRecusiveMem(num int) int {
	if result, ok := resultMap[num]; ok {
		return result
	}
	if num == 1 || num == 2 {
		return 1
	}
	result := fibonacciRecusive(num-1) + fibonacciRecusive(num-2)
	resultMap[num] = result
	return result
}

// 遍历的方式 - 时间复杂度 O(n)
func fibonacciIterative(num int) int {
	var f = []int{
		0, 1,
	}

	for i := 2; i <= num; i++ {
		f = append(f, f[i-1]+f[i-2])
	}

	return f[num]
}

// todo 使用二维矩阵的方式来将时间复杂度优化到 O(logn)
func fibonacciMaxtrix(num int) int {
	return 0
}

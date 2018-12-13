package matrix

// matrix 输入一个整数n, 输出一个n*n的矩阵

func matrix(num int) (metrixs [][]int) {
	for i := 0; i < num; i++ {
		var line []int
		for j := 1; j <= num; j++ {
			line = append(line, i*num+j)
		}
		metrixs = append(metrixs, line)
	}
	return
}

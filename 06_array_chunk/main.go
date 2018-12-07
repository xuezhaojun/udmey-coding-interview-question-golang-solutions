package main

import (
	"fmt"
)

// array Chunk 用于将一个长的数组，拆分为多个小的数组

func main() {
	var intArray []int = []int{
		1, 2, 3, 4, 5,
	}
	intGroup := arrayChunk(intArray, 2)
	for _, intarray := range intGroup {
		fmt.Printf("%#v\n", intarray)
	}
	intGroup2 := arrayChunk2(intArray, 2)
	for _, intarray := range intGroup2 {
		fmt.Printf("%#v\n", intarray)
	}
}

// solution_1 先组成一个子array，再添加到整个group中
func arrayChunk(intarray []int, chunk int) (result [][]int) {
	arrayLength := len(intarray)
	tempArray := []int{}
	for i := 1; i <= arrayLength; i++ {
		tempArray = append(tempArray, i)
		if i%chunk == 0 {
			result = append(result, tempArray)
			tempArray = []int{}
		}
	}
	result = append(result, tempArray)
	return
}

// solution_better 先给整个的group添加array，再像 array 中添加item
func arrayChunk2(intarray []int, chunk int) (result [][]int) {
	for i, num := range intarray {
		if i%chunk == 0 || len(result) == 0 { // 只有当result为空，或者是达到要增加array的上限时，才增加一个array
			result = append(result, []int{})
		}
		result[len(result)-1] = append(result[len(result)-1], num) // 给 result 的最后一个array添加 item
	}
	return
}

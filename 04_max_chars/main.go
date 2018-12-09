package main

//  max chars : 找到一个字符串中，数量最多的那一个字符
//  遇到字符串相关的算法，首先要问自己3个问题：
// 1. 字符串中最常见的字符是什么？
// 2. 输入和输出的字符串有相同的字符吗？
// 3. 给出的字符串中有重复的字符吗？

// 使用map
func maxChars(str string) string {
	bytesMap := make(map[byte]int)
	bs := []byte(str)
	if len(bs) == 0 {
		return ""
	}

	var maxchar byte = 0
	maxcount := 0
	for _, b := range bs {
		count := bytesMap[b]
		bytesMap[b] = count + 1
		// 判断当前的字符是否大于当前最大
		if bytesMap[b] > maxcount {
			maxchar = b
			maxcount = bytesMap[b]
		}
	}

	return string([]byte{maxchar})
}

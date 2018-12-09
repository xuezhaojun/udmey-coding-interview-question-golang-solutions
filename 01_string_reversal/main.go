package main

// 遍历字符串
// golang 中没有build-in的例如js中的 reverse()的方法对数组进行反转，所以得手写一个反转数组的方法
// 为什么这里使用了 rune ,而不是 byte ?
// byte 是 8bit
// rune 是 32bit
// 在utf-8编码中，对于 中文字符 而言，一个字符占 3个字节, 使用 byte 是放不下的
// 常见的 range 也是对str进行了隐式的 unicode 解码, 而 str[i] 并不一定和我们看到的字符串对应
// 同理，如果只是序列化和反序列化，可以通过byte进行操作，但是如果涉及字符串中的反转，截断等操作，则使用rune
func reverse02(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 使用byte进行反转字符串的话，仅会对字母和数字有效，一旦加入中文字符串，结果就无法满足预期了
func reverseByByte(str string) string {
	bs := []byte(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
	return string(bs)
}

package _1_vowels

import "strings"

// 获取一句话中所有的元音字母的数量
// 最简单的方式就是遍历字符串中的所有字节，对每一个进行判断是否在a,e,i,o,u之中
// 复杂的方式为，通过正则表达式取匹配

func vowels(str string) (num int) {
	var vowelsmap = make(map[rune]struct{})
	vowelsmap['a'] = struct{}{}
	vowelsmap['e'] = struct{}{}
	vowelsmap['i'] = struct{}{}
	vowelsmap['o'] = struct{}{}
	vowelsmap['u'] = struct{}{}
	str = strings.ToLower(str)
	for _, s := range str {
		if _, ok := vowelsmap[s]; ok {
			num++
		}
	}
	return
}

package main

import (
	"sort"
	"strings"
)

// anagram 是指，两个字符串中，如果字符的数量相同，则是 anagram 如果不同，则不是
// 对应leetCode 242

// 使用map对str1和str2中的不同字节进行计数
// 对 str1 进行加
// 对 str2 进行减
func anagram(s, t string) (result bool) {
	m := make(map[rune]int)
	runes1 := []rune(s)
	runes2 := []rune(t)

	if len(runes1) != len(runes2) {
		return false
	}

	length := len(runes1)
	for i := 0; i < length; i++ {
		m[runes1[i]] += 1
		m[runes2[i]] -= 1
	}

	for _, c := range m {
		if c != 0 {
			return false
		}
	}

	return true
}

// anagram 等同于 排序后完全一样
func anagram2(s, t string) bool {
	return stringSort(s) == stringSort(t)
}

func stringSort(s string) string {
	ss := strings.Split(s, "")
	sort.Strings(ss)
	return strings.Join(ss, "")
}

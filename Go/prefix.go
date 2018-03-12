package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := []string{
		"ab",
		"ab",
		"abc",
	}
	fmt.Println(longestCommonPrefix(s))
}

func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}

	// 记录最长公共字符串前缀
	var b bytes.Buffer
	// 记录每一次遍历的公共字符串
	var ch string
	// 外层记录公共字符前缀的最大长度
	for i := 0; ; i++ {
		ch = ""
		// 内层每次均需要遍历整个字符数组
		// 依次寻找第一个前缀, 第二个前缀...
		for j := 0; j < length; j++ {
			// 一个子字符串的长度, 当前遍历到最长的子字符串
			if i >= len(strs[j]) {
				return b.String()
			}
			if ch == "" {
				// 存储一个一个的字符
				ch = string(strs[j][i])
				// 即公共的子字符串到此为止
			} else if string(strs[j][i]) != ch {
				return b.String()
			}
		}
		// 每次记录一个公共前缀
		b.WriteString(ch)
	}
}

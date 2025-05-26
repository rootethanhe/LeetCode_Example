package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	// 预处理空格并分割单词(自动处理所有多余空格)
	words := strings.Fields(s)

	// 反转单词数组(双指针法)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// 重组字符串(单词间单个空格)
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(reverseWords("the sky is blue"))  //"blue is sky the"
	fmt.Println(reverseWords("  hello world  "))  //"world hello"
	fmt.Println(reverseWords("a good   example")) //"example good a"
}

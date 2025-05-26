package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	sList := strings.Split(s, " ")
	if len(pattern) != len(sList) { // 长度不同直接返回false
		return false
	}

	pMap := make(map[byte]string) // 记录p→s的映射
	sMap := make(map[string]byte) // 记录s→p的映射

	for i := 0; i < len(pattern); i++ {
		pChar := pattern[i]
		sString := sList[i]

		// 检查p→s的映射是否冲突
		if mappedPChar, exists := sMap[sString]; exists {
			if mappedPChar != pChar {
				return false
			}
		} else {
			// 检查s→p的映射是否冲突
			if mappedSString, exists := pMap[pChar]; exists {
				if mappedSString != sString {
					return false
				}
			}
		}

		// 更新双向映射
		sMap[sString] = pChar
		pMap[pChar] = sString
	}

	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))  // true
	fmt.Println(wordPattern("abba", "dog cat cat fish")) // false
	fmt.Println(wordPattern("aaaa", "dog cat cat dog"))  // false
	fmt.Println(wordPattern("abba", "dog dog dog dog"))  // false
}

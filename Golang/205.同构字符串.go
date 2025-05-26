package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) { // 长度不同直接返回false
		return false
	}

	sMap := make(map[byte]byte) // 记录s→t的映射
	tMap := make(map[byte]byte) // 记录t→s的映射

	for i := 0; i < len(s); i++ {
		sChar := s[i]
		tChar := t[i]

		// 检查s→t的映射是否冲突
		if mappedTChar, exists := sMap[sChar]; exists {
			if mappedTChar != tChar {
				return false
			}
		} else {
			// 检查t→s的映射是否冲突
			if mappedSChar, exists := tMap[tChar]; exists {
				if mappedSChar != sChar {
					return false
				}
			}
			// 更新双向映射
			sMap[sChar] = tChar
			tMap[tChar] = sChar
		}
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("egg", "add"))     // true（e→a, g→d）
	fmt.Println(isIsomorphic("foo", "bar"))     // false（o无法同时映射到a和r）
	fmt.Println(isIsomorphic("paper", "title")) // true（p→t, a→i, e→l, r→e）
	fmt.Println(isIsomorphic("ab", "aa"))       // false（b→a与a→a冲突）
}

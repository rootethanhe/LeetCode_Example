package main

import "fmt"

func isAnagram(s string, t string) bool {
	// 长度不同直接返回false
	if len(s) != len(t) {
		return false
	}

	// 初始化计数数组(26个小写字母)
	counts := [26]int{}

	// 遍历s和t，更新计数
	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']++ // s的字符计数+1
		counts[t[i]-'a']-- // t的字符计数-1
	}

	// 检查所有计数是否为0
	for _, c := range counts {
		if c != 0 {
			return false
		}
	}

	return true
}

func main() {
	// 示例测试
	fmt.Println(isAnagram("anagram", "nagaram")) // true
	fmt.Println(isAnagram("rat", "car"))         // false
	fmt.Println(isAnagram("listen", "silent"))   // true
	fmt.Println(isAnagram("hello", "world"))     // false
}

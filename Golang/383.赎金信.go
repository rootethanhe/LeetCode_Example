package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	// 长度不足时直接失败
	if len(ransomNote) > len(magazine) {
		return false
	}

	// 初始化计数数组(26个小写字母)
	counts := [26]int{}

	// 统计magazine字符出现次数
	for _, ch := range magazine {
		counts[ch-'a']++ // 'a'-->0, 'b'-->1, ..., 'z'-->25
	}

	// 验证ransomNote字符可用性
	for _, ch := range ransomNote {
		index := ch - 'a'
		counts[index]--        // 消耗字符计数
		if counts[index] < 0 { // 字符数量不足
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(canConstruct("a", "b"))   // false
	fmt.Println(canConstruct("aa", "ab")) // false
	fmt.Println(canConstruct("a", "aab")) // true
}

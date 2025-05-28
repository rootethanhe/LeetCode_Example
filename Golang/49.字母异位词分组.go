package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	// 创建哈希表: 键=排序后字符串，值=原始字符串列表
	groups := make(map[string][]string)

	for _, str := range strs {
		// 1. 字符串转字符切片
		chars := strings.Split(str, "")

		// 2. 对字符切片进行排序（标准化）
		sort.Strings(chars)

		// 3. 生成排序后的键
		key := strings.Join(chars, "")

		// 4. 将原始字符串加入对应分组
		groups[key] = append(groups[key], str)
	}

	// 5. 收集结果
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

func main() {
	// 测试用例
	cases := [][]string{
		{"eat", "tea", "tan", "ate", "nat", "bat"}, // 标准输入
		{""},  // 空字符串
		{"a"}, // 单字符
	}

	for _, strs := range cases {
		fmt.Printf("输入: %v\n输出: %v\n\n", strs, groupAnagrams(strs))
	}
}

package main

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	// 数字到字母的映射表
	phoneMap := []string{
		"",     // 0
		"",     // 1
		"abc",  // 2
		"def",  // 3
		"ghi",  // 4
		"jkl",  // 5
		"mno",  // 6
		"pqrs", // 7
		"tuv",  // 8
		"wxyz", // 9
	}

	res := []string{}
	backtrack(&res, phoneMap, digits, 0, []byte{})
	return res
}

func backtrack(res *[]string, phoneMap []string, digits string, index int, path []byte) {
	// 终止条件：完成所有数字的处理
	if index == len(digits) {
		*res = append(*res, string(path))
		return
	}

	// 获取当前数字对应的字母集合
	digit := digits[index] - '0'
	letters := phoneMap[digit]

	// 遍历所有可能性
	for i := 0; i < len(letters); i++ {
		path = append(path, letters[i])
		backtrack(res, phoneMap, digits, index+1, path)
		path = path[:len(path)-1] // 回溯
	}
}

func main() {
	// 示例1：输入"23"
	fmt.Println(letterCombinations("23"))
	// 输出：[ad ae af bd be bf cd ce cf]

	// 示例2：输入空字符串
	fmt.Println(letterCombinations(""))
	// 输出：[]

	// 示例3：输入"2"
	fmt.Println(letterCombinations("2"))
	// 输出：[a b c]
}

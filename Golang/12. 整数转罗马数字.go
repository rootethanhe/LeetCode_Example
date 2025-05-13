package main

import "fmt"

func intToRoman(num int) string {
	// 预定义符号映射表（按数值降序排列）
	symbols := []struct {
		value  int
		symbol string
	}{
		{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"},
		{1, "I"},
	}

	roman := ""
	// 贪心算法：优先选择最大可减的符号
	for _, s := range symbols {
		for num >= s.value {
			roman += s.symbol
			num -= s.value
		}
		if num == 0 {
			break
		}
	}
	return roman
}

func main() {
	// 示例1
	num1 := 3749
	fmt.Println(intToRoman(num1)) // 输出: "MMMDCCXLIX"

	// 示例2
	num2 := 1994
	fmt.Println(intToRoman(num2)) // 输出: "MCMXCIV"
}

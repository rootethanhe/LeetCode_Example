package main

import "fmt"

func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	carry := 0
	result := make([]byte, 0)

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		// 获取当前位的数值（如果存在）
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		// 计算当前位结果和进位
		carry = sum / 2
		result = append(result, byte(sum%2+'0'))
	}
	// 反转结果字符串
	reverse(result)
	return string(result)
}

// 反转字节数组
func reverse(bytes []byte) {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
}

func main() {
	fmt.Println(addBinary("11", "1"))      // 输出 "100"
	fmt.Println(addBinary("1010", "1011")) // 输出 "10101"
	fmt.Println(addBinary("1", "111"))     // 输出 "1000"
}

package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) != len(cost) {
		return -1
	}
	totalGas := 0   // 总油量剩余
	currentGas := 0 // 当前累积油量
	start := 0      // 可能的起点索引

	for i := 0; i < len(gas); i++ {
		totalGas += gas[i] - cost[i]
		currentGas += gas[i] - cost[i]
		// 当前累积油量为负，说明起点不在当前及之前的加油站
		if currentGas < 0 {
			start = i + 1  // 重置起点为下一加油站
			currentGas = 0 // 重置当前油量
		}
	}

	// 总油量不足则无解，否则返回起点
	if totalGas >= 0 && start < len(gas) {
		return start
	}
	return -1
}

func main() {
	// 示例1
	gas1 := []int{1, 2, 3, 4, 5}
	cost1 := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas1, cost1)) // 输出: 3

	// 示例2
	gas2 := []int{2, 3, 4}
	cost2 := []int{3, 4, 3}
	fmt.Println(canCompleteCircuit(gas2, cost2)) // 输出: -1
}

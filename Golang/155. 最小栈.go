package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	data     []int // 数据栈
	minStack []int // 最小值栈
}

func Constructor() MinStack {
	return MinStack{
		data:     make([]int, 0),
		minStack: []int{math.MaxInt32}, // 初始化为极大值，避免空栈判断
	}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	// 计算当前最小值（新值与 minStack 栈顶的较小者）
	currentMin := min(val, this.minStack[len(this.minStack)-1])
	this.minStack = append(this.minStack, currentMin)
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

// 辅助函数：取较小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println("当前最小值:", obj.GetMin()) // 输出: -3
	obj.Pop()
	fmt.Println("栈顶元素:", obj.Top())     // 输出: 0
	fmt.Println("当前最小值:", obj.GetMin()) // 输出: -2
}

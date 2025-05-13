package main

import (
	"fmt"
	"math/rand"
	"time"
)

type RandomizedSet struct {
	data  []int       // 存储元素的动态数组
	index map[int]int // 元素到数组索引的映射
}

func Constructor() RandomizedSet {
	rand.Seed(time.Now().UnixNano()) // 初始化随机种子
	return RandomizedSet{
		data:  make([]int, 0),
		index: make(map[int]int),
	}
}

// Insert 插入元素(若不存在)
func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.index[val]; exists {
		return false
	}
	this.data = append(this.data, val)   // 追加到数组末尾
	this.index[val] = len(this.data) - 1 // 记录索引
	return true
}

// Remove 删除元素(若存在)
func (this *RandomizedSet) Remove(val int) bool {
	idx, exists := this.index[val]
	if !exists {
		return false
	}

	lastIdx := len(this.data) - 1
	lastVal := this.data[lastIdx]

	// 将最后一个元素移动到待删除位置
	this.data[idx] = lastVal
	this.index[lastVal] = idx

	// 删除原元素
	this.data = this.data[:lastIdx]
	delete(this.index, val)

	return true
}

// GetRandom 返回随机元素
func (this *RandomizedSet) GetRandom() int {
	if len(this.data) == 0 {
		return -1 // 根据题目约束，调用时必有元素
	}
	return this.data[rand.Intn(len(this.data))]
}

func main() {
	rs := Constructor()
	fmt.Println(rs.Insert(1))   // true
	fmt.Println(rs.Remove(2))   // false
	fmt.Println(rs.Insert(2))   // true
	fmt.Println(rs.GetRandom()) // 1 或 2
	fmt.Println(rs.Remove(1))   // true
	fmt.Println(rs.Insert(2))   // false
	fmt.Println(rs.GetRandom()) // 2
}

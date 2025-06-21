package main

import "fmt"

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	n := len(grid)
	return build(0, 0, n, grid)
}

// 递归构建四叉树：r0/c0=起始行列，size=当前区域边长
func build(r0, c0, size int, grid [][]int) *Node {
	// 检查当前区域是否全相同
	allSame, val := checkUniform(r0, c0, size, grid)
	if allSame {
		return &Node{Val: val, IsLeaf: true} // 叶子节点
	}

	// 非叶子节点：划分四个子区域
	half := size / 2
	return &Node{
		IsLeaf:      false,
		Val:         true,                                // 任意值，判题机制接受
		TopLeft:     build(r0, c0, half, grid),           // 左上
		TopRight:    build(r0, c0+half, half, grid),      // 右上
		BottomLeft:  build(r0+half, c0, half, grid),      // 左下
		BottomRight: build(r0+half, c0+half, half, grid), // 右下
	}
}

// 检查区域值是否一致
func checkUniform(r0, c0, size int, grid [][]int) (bool, bool) {
	base := grid[r0][c0]
	for i := r0; i < r0+size; i++ {
		for j := c0; j < c0+size; j++ {
			if grid[i][j] != base {
				return false, false // 存在不同值
			}
		}
	}
	return true, base == 1 // 全部相同，返回true和值（true对应1）
}

// 打印四叉树, 测试使用
func printTree(root *Node) {
	if root == nil {
		fmt.Println("[]")
		return
	}
	queue := []*Node{root}
	var res [][]int

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			res = append(res, nil)
		} else {
			res = append(res, []int{boolToInt(node.IsLeaf), boolToInt(node.Val)})
			queue = append(queue, node.TopLeft, node.TopRight, node.BottomLeft, node.BottomRight)
		}
	}

	// 移除末尾 nil
	i := len(res) - 1
	for i >= 0 && res[i] == nil {
		i--
	}
	fmt.Println(res[:i+1])
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func main() {
	// 示例1：2x2矩阵
	grid1 := [][]int{{0, 1}, {1, 0}}
	root1 := construct(grid1)
	printTree(root1) // 输出序列化格式（需自行实现层序遍历）

	// 示例2：8x8矩阵
	grid2 := [][]int{
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
	}
	root2 := construct(grid2)
	printTree(root2)

	// 边界测试：全1矩阵
	grid3 := [][]int{{1, 1}, {1, 1}}
	root3 := construct(grid3) // 应返回单个叶子节点 [1,1]
	printTree(root3)
}

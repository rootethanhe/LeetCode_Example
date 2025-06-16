package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])

	// 起点或终点是障碍物, 直接返回0
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	// 初始化一维DP数组
	dp := make([]int, n)
	dp[0] = 1 // 起点路径数为1

	// 初始化首行: 从左向右扫描
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			dp[j] = 0 // 障碍物位置路径数为0
		} else {
			dp[j] = dp[j-1] // 继承左侧路径数
		}
	}

	// 从第二行开始逐行更新
	for i := 1; i < m; i++ {
		// 更新当前行首列
		if obstacleGrid[i][0] == 1 {
			dp[0] = 0 // 障碍物位置路径数为0
		}
		// 更新当前行其他列
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0 // 障碍物位置路径数为0
			} else {
				dp[j] = dp[j] + dp[j-1] // 状态转移: 上方路径数 + 左方路径数
			}
		}
	}

	return dp[n-1]
}

func main() {
	// 示例1：3x3网格，中心障碍物
	grid1 := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println(uniquePathsWithObstacles(grid1)) // 输出：2

	// 示例2：2x2网格，右上角障碍物
	grid2 := [][]int{{0, 1}, {0, 0}}
	fmt.Println(uniquePathsWithObstacles(grid2)) // 输出：1

	// 边界测试：1x1网格无障碍
	grid3 := [][]int{{0}}
	fmt.Println(uniquePathsWithObstacles(grid3)) // 输出：1

	// 边界测试：1x1网格有障碍
	grid4 := [][]int{{1}}
	fmt.Println(uniquePathsWithObstacles(grid4)) // 输出：0

	// 障碍物阻断路径
	grid5 := [][]int{{0, 0, 0}, {0, 1, 1}, {0, 1, 0}}
	fmt.Println(uniquePathsWithObstacles(grid5)) // 输出：0
}

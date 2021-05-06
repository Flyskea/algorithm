/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 */
package main

import "fmt"

// @lc code=start

var dx []int = []int{0, 1, -1, 0}
var dy []int = []int{1, 0, 0, -1}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(grid [][]int, row, col int) int {
	rows := len(grid)
	cols := len(grid[0])
	if grid[row][col] == 0 {
		return 0
	}
	grid[row][col] = 0
	var x, y, area int
	area = 1
	for i := 0; i < 4; i++ {
		x = row + dx[i]
		y = col + dy[i]
		if x >= 0 && x < rows && y >= 0 && y < cols {
			area += dfs(grid, x, y)
		}
	}
	return area
}

func maxAreaOfIsland(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	var maxArea int
	if rows == 0 || cols == 0 {
		return 0
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				maxArea = max(maxArea, dfs(grid, i, j))
			}
		}
	}
	return maxArea
}

// @lc code=end

func main() {
	fmt.Println(maxAreaOfIsland([][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}}))
}

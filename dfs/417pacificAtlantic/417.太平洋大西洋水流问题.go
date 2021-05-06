/*
 * @lc app=leetcode.cn id=417 lang=golang
 *
 * [417] 太平洋大西洋水流问题
 */
package main

import (
	"fmt"
)

// @lc code=start

var dx []int = []int{0, 1, -1, 0}
var dy []int = []int{1, 0, 0, -1}

func dfs(heights [][]int, row, col int, visited *[][]bool) {
	if (*visited)[row][col] {
		return
	}
	var x, y int
	(*visited)[row][col] = true
	for i := 0; i < 4; i++ {
		x = row + dx[i]
		y = col + dy[i]
		if x >= 0 && x < len(heights) &&
			y >= 0 && y < len(heights[0]) && heights[x][y] >= heights[row][col] {
			dfs(heights, x, y, visited)
		}
	}
}

func pacificAtlantic(heights [][]int) [][]int {
	rows, cols := len(heights), len(heights[0])
	if rows == 0 || cols == 0 {
		return nil
	}
	var res [][]int
	pacific, atlantic := make([][]bool, rows), make([][]bool, rows)
	for i := 0; i < rows; i++ {
		pacific[i] = make([]bool, cols)
		atlantic[i] = make([]bool, cols)
	}
	for i := 0; i < rows; i++ {
		dfs(heights, i, 0, &pacific)
		dfs(heights, i, cols-1, &atlantic)
	}
	for j := 0; j < cols; j++ {
		dfs(heights, 0, j, &pacific)
		dfs(heights, rows-1, j, &atlantic)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if atlantic[i][j] && pacific[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

// @lc code=end
func main() {
	fmt.Println(pacificAtlantic([][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}))
}

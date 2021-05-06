/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 省份数量
 */
package main

import "fmt"

// @lc code=start
func dfs(isConnected [][]int, cur int, visited []bool) {
	visited[cur] = true
	cols := len(isConnected[cur])
	for j := 0; j < cols; j++ {
		if !visited[j] && isConnected[cur][j] == 1 {
			dfs(isConnected, j, visited)
		}
	}
}
func findCircleNum(isConnected [][]int) int {
	rows := len(isConnected)
	if rows <= 0 {
		return 0
	}
	res := 0
	visited := make([]bool, rows)
	for i := range isConnected {
		if !visited[i] {
			dfs(isConnected, i, visited)
			res++
		}
	}
	return res
}

// @lc code=end

func main() {
	fmt.Println(
		findCircleNum([][]int{{1, 1, 0},
			{1, 1, 0}, {0, 0, 1}},
		))
}

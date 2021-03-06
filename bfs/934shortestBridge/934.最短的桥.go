/*
 * @lc app=leetcode.cn id=934 lang=golang
 *
 * [934] 最短的桥
 */
package main

import (
	"container/list"
	"fmt"
)

// @lc code=start

type node struct {
	x    int
	y    int
	step int
}

var gx = []int{-1, 1, 0, 0}
var gy = []int{0, 0, -1, 1}

func shortestBridge(A [][]int) int {
	n := len(A)
	//标记第一个找到的岛屿
	vis := make([][]bool, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]bool, n)
	}

	// 将岛屿边界入队列
	que := list.New()
	setVis(A, vis, que)

	// bfs求最短路径
	for que.Len() > 0 {
		qf := que.Front()
		u := qf.Value.(node)
		// 第一次找到A[][]=1且没有被访问的即时最短路径
		if !vis[u.x][u.y] && A[u.x][u.y] == 1 {
			return u.step - 1
		}
		que.Remove(qf)
		//遍历4个方向
		for i := 0; i < 4; i++ {
			xx := u.x + gx[i]
			yy := u.y + gy[i]

			if xx < 0 || xx >= n || yy < 0 || yy >= n {
				continue
			}

			if vis[xx][yy] {
				continue
			}
			if !vis[xx][yy] && A[xx][yy] == 0 {
				vis[xx][yy] = true
			}
			que.PushBack(node{xx, yy, u.step + 1})
		}
	}
	return -1
}

// dfs标记第一个岛屿
func setVis(grid [][]int, vis [][]bool, que *list.List) {
	n := len(grid)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dfs(grid, i, j, vis, que)
				return
			}
		}
	}
}

// 标记第一个岛屿
func dfs(grid [][]int, i int, j int, vis [][]bool, que *list.List) {
	n := len(grid)
	if i < 0 || i >= n || j < 0 || j >= n {
		return
	}
	if vis[i][j] {
		return
	}
	if grid[i][j] == 0 {
		return
	}
	if !vis[i][j] && grid[i][j] == 1 {
		vis[i][j] = true
	}
	// 将边界写入队列
	if isEdge(grid, i, j) {
		que.PushBack(node{i, j, 0})
		//return
	}
	//遍历上下左右
	dfs(grid, i-1, j, vis, que)
	dfs(grid, i+1, j, vis, que)
	dfs(grid, i, j-1, vis, que)
	dfs(grid, i, j+1, vis, que)
}

func isEdge(grid [][]int, i int, j int) bool {
	n := len(grid)
	if i-1 >= 0 && grid[i-1][j] == 0 {
		return true
	}
	if i+1 < n && grid[i+1][j] == 0 {
		return true
	}
	if j-1 >= 0 && grid[i][j-1] == 0 {
		return true
	}
	if j+1 < n && grid[i][j+1] == 0 {
		return true
	}
	return false
}

// @lc code=end

func main() {
	fmt.Println(shortestBridge([][]int{{0, 1}, {1, 0}}))
}

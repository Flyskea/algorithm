/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 */
package main

import "fmt"

// @lc code=start
var dx []int = []int{0, -1, 1, 0}
var dy []int = []int{1, 0, 0, -1}

func solve(board [][]byte) {
	for i := range board {
		for j := range board[i] {
			if i == 0 || i == len(board)-1 || j == 0 || j == len(board[i])-1 {
				if board[i][j] == 'O' {
					dfs(i, j, board)
				}
			}
		}
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == '*' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(row, col int, board [][]byte) {
	if board[row][col] == 'O' {
		board[row][col] = '*'
		var x, y int
		for i := 0; i < 4; i++ {
			x, y = row+dx[i], col+dy[i]
			if x >= 0 && x < len(board) && y >= 0 && y < len(board[0]) && board[x][y] != 'X' {
				dfs(x, y, board)
			}
		}
	}
}

// @lc code=end

// BFS
func solve1(board [][]byte) {
	rows, cols := len(board), len(board[0])
	if rows < 1 || cols < 1 {
		return
	}
	var queue [][]int
	for i := range board {
		for j := range board[i] {
			if (i == 0 || j == 0 || i == rows-1 || j == cols-1) && board[i][j] == 'O' {
				queue = append(queue, []int{i, j})
			}
		}
	}
	for len(queue) > 0 {
		queueLen := len(queue)
		for i := 0; i < queueLen; i++ {
			node := queue[0]
			queue = queue[1:]
			row, col := node[0], node[1]
			if board[row][col] == '*' {
				continue
			}
			board[row][col] = '*'
			var x, y int
			for j := 0; j < 4; j++ {
				x, y = row+dx[j], col+dy[j]
				if x >= 0 && y >= 0 && x < rows && y < cols && board[x][y] != 'X' {
					queue = append(queue, []int{x, y})
				}
			}
		}
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == '*' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func main() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve1(board)
	fmt.Println(board)
}

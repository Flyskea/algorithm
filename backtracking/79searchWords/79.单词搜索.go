/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */
package main

import "fmt"

// @lc code=start
var dx []int = []int{-1, 0, 0, 1}
var dy []int = []int{0, -1, 1, 0}

func backtracking(board [][]byte, visited [][]bool, word string, index, row, col int) bool {
	if index == len(word)-1 {
		return board[row][col] == word[index]
	}
	var x, y int
	if board[row][col] == word[index] {
		visited[row][col] = true
		for i := 0; i < 4; i++ {
			x = row + dx[i]
			y = col + dy[i]
			if x >= 0 && x < len(board) && y >= 0 && y < len(board[0]) && !visited[x][y] && backtracking(board, visited, word, index+1, x, y) {
				return true
			}
		}
		visited[row][col] = false
	}
	return false
}

func exist(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}
	for i, v := range board {
		for j := range v {
			if backtracking(board, visited, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

// @lc code=end

func main() {
	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "SEE"))
}

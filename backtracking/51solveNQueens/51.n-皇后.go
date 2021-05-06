/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */
package main

import (
	"fmt"
	"strings"
)

// @lc code=start
func backtracking(row int, board [][]string, res *[][]string, n int, cols, diag1, diag2 map[int]bool) {
	if row == n {
		temp := make([]string, len(board))
		for i := 0; i < n; i++ {
			temp[i] = strings.Join(board[i], "")
		}
		*res = append(*res, temp)
		return
	}
	for col := 0; col < n; col++ {
		if !cols[col] && !diag1[row+col] && !diag2[row-col] {
			board[row][col] = "Q"
			cols[col] = true
			diag1[row+col] = true
			diag2[row-col] = true
			backtracking(row+1, board, res, n, cols, diag1, diag2)
			board[row][col] = "."
			cols[col] = false
			diag1[row+col] = false
			diag2[row-col] = false
		}
	}
}

func solveNQueens(n int) [][]string {
	bd := make([][]string, n)
	for i := range bd {
		bd[i] = make([]string, n)
		for j := range bd[i] {
			bd[i][j] = "."
		}
	}
	cols := map[int]bool{}
	diag1 := map[int]bool{}
	diag2 := map[int]bool{}

	res := [][]string{}
	backtracking(0, bd, &res, n, cols, diag1, diag2)
	return res
}

// @lc code=end

func main() {
	fmt.Println(solveNQueens(4))
}

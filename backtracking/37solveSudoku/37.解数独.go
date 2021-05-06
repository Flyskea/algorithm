/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */
package main

// @lc code=start
func hasConflict(board [][]byte, row, col int, val byte) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == val || board[row][i] == val {
			return true
		}
	}
	subRowStart := (row / 3) * 3
	subColStart := (col / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[subRowStart+i][subColStart+j] == val {
				return true
			}
		}
	}
	return false
}

func backtracking(board [][]byte, row, col int, rows, cols, blocks [9]map[byte]bool) bool {
	if col == 9 {
		row++
		col = 0
		if row == 9 {
			return true
		}
	}
	if board[row][col] != '.' {
		return backtracking(board, row, col+1, rows, cols, blocks)
	}
	blockIndex := getBlockIndex(row, col)
	for i := byte('1'); i <= byte('9'); i++ {
		if !rows[row][i] || !cols[col][i] || !blocks[blockIndex][i] {
			continue
		}
		board[row][col] = i
		rows[row][i] = false
		cols[col][i] = false
		blocks[blockIndex][i] = false
		if backtracking(board, row, col+1, rows, cols, blocks) {
			return true
		}
		rows[row][i] = true
		cols[col][i] = true
		blocks[blockIndex][i] = true
		board[row][col] = '.'
	}
	return false
}

func getBlockIndex(i, j int) int {
	return (i/3)*3 + j/3
}

func solveSudoku(board [][]byte) {
	rows := [9]map[byte]bool{}
	cols := [9]map[byte]bool{}
	blocks := [9]map[byte]bool{}
	for i := 0; i < 9; i++ {
		rows[i] = map[byte]bool{'1': true, '2': true, '3': true, '4': true, '5': true, '6': true, '7': true, '8': true, '9': true}
		cols[i] = map[byte]bool{'1': true, '2': true, '3': true, '4': true, '5': true, '6': true, '7': true, '8': true, '9': true}
		blocks[i] = map[byte]bool{'1': true, '2': true, '3': true, '4': true, '5': true, '6': true, '7': true, '8': true, '9': true}
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				rows[i][board[i][j]] = false
				cols[j][board[i][j]] = false
				blocks[getBlockIndex(i, j)][board[i][j]] = false
			}
		}
	}
	backtracking(board, 0, 0, rows, cols, blocks)
}

// @lc code=end

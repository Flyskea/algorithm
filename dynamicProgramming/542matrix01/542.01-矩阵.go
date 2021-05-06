/*
 * @lc app=leetcode.cn id=542 lang=golang
 *
 * [542] 01 矩阵
 */
package main

import (
	"fmt"
	"math"
)

// @lc code=start
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func updateMatrix(matrix [][]int) [][]int {
	rows, cols := len(matrix), len(matrix[0])
	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
		for k := range dp[i] {
			dp[i][k] = math.MaxInt32
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				dp[i][j] = 0
			} else {
				if j > 0 {
					dp[i][j] = min(dp[i][j], dp[i][j-1]+1)
				}
				if i > 0 {
					dp[i][j] = min(dp[i][j], dp[i-1][j]+1)
				}
			}
		}
	}
	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 0; j-- {
			if matrix[i][j] != 0 {
				if j < cols-1 {
					dp[i][j] = min(dp[i][j], dp[i][j+1]+1)
				}
				if i < rows-1 {
					dp[i][j] = min(dp[i][j], dp[i+1][j]+1)
				}
			}
		}
	}
	return dp
}

// @lc code=end

func main() {
	fmt.Println(updateMatrix([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}))
}

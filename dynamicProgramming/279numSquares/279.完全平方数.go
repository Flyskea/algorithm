/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
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

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

// @lc code=end

func main() {
	fmt.Println(numSquares(2))
}

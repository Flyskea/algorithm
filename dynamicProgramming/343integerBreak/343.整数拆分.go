/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 */
package main

import (
	"fmt"
)

// @lc code=start

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func integerBreak(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		curMax := 0
		for j := 1; j < i; j++ {
			curMax = max(curMax, max(j*(i-j), j*dp[i-j]))
		}
		dp[i] = curMax
	}
	res := dp[n]
	return res
}

// @lc code=end

func main() {
	fmt.Println(integerBreak(10))
}

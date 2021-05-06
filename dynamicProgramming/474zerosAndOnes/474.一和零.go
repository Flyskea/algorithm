/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 */
package main

import (
	"fmt"
	"strings"
)

// @lc code=start

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for _, s := range strs {
		zero := strings.Count(s, "0")
		one := len(s) - zero
		if zero > m || one > n {
			continue
		}
		for i := m; i >= zero; i-- {
			for j := n; j >= one; j-- {
				dp[i][j] = max(dp[i][j], 1+dp[i-zero][j-one])
			}
		}
	}
	return dp[m][n]
}

// @lc code=end

func main() {
	fmt.Println(findMaxForm([]string{"10", "0001",
		"111001", "1", "0"}, 5, 3))
}

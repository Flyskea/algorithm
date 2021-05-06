/*
 * @lc app=leetcode.cn id=650 lang=golang
 *
 * [650] 只有两个键的键盘
 */
package main

import (
	"fmt"
)

// @lc code=start
func minSteps(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = i
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				dp[i] = dp[j] + dp[i/j]
				break
			}
		}
	}
	return dp[n]
}

// @lc code=end

func main() {
	fmt.Println(minSteps(6))
}

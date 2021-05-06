/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */
package main

import "fmt"

// @lc code=start

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxProfit(k int, prices []int) int {
	if k < 1 || len(prices) < 1 {
		return 0
	}
	dp := make([][2]int, k+1)

	for i := 0; i < len(prices); i++ {
		for j := k; j > 0; j-- {
			if i == 0 {
				dp[j-1][1] = -prices[i]
				dp[j-1][0] = 0
			} else {
				dp[j-1][1] = max(dp[j-1][1], dp[j-1][0]-prices[i])
				dp[j][0] = max(dp[j][0], dp[j-1][1]+prices[i])
			}
		}
	}
	return dp[k][0]
}

// @lc code=end

func main() {
	fmt.Println(maxProfit(2, []int{2}))
}

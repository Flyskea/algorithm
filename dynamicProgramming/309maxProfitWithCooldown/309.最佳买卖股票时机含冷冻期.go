/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
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

func maxProfit(prices []int) int {
	pricesLen := len(prices)
	if pricesLen < 1 {
		return 0
	}
	dp := make([][3]int, pricesLen)
	dp[0][0] = -prices[0]
	for i := 1; i < pricesLen; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	res := max(dp[pricesLen-1][1], dp[pricesLen-1][2])
	return res
}

// @lc code=end

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
}

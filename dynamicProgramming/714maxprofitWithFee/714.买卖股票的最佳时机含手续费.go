/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
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

func maxProfit(prices []int, fee int) int {
	pricesLen := len(prices)
	if pricesLen < 1 {
		return 0
	}
	dp := make([][2]int, pricesLen)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < pricesLen; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[pricesLen-1][0]
}

// @lc code=end

func main() {
	fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
}

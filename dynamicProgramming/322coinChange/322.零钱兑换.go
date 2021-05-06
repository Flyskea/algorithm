/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */
package main

import "fmt"

// @lc code=start
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	for i := range coins {
		coin := coins[i]
		for j := coins[i]; j <= amount; j++ {
			dp[j] = min(dp[j], dp[j-coin]+1)
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// @lc code=end

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}

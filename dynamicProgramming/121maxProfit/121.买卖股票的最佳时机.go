/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */
package main

import (
	"fmt"
)

// @lc code=start
func maxProfit(prices []int) int {
	minValue := prices[0]
	maxValue := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minValue {
			minValue = prices[i]
		}
		if prices[i]-minValue > maxValue {
			maxValue = prices[i] - minValue
		}
	}
	return maxValue
}

// @lc code=end

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6}))
}

/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */
package main

import "fmt"

// 贪婪解法 不同写法 一直寻找最大的从而局部最优
// @lc code=start

func maxProfit(prices []int) int {
	pricesLen := len(prices)
	var profit int
	var res int
	for i := 1; i < pricesLen; i++ {
		profit = prices[i] - prices[i-1]
		if profit > 0 {
			res += profit
		}
	}
	return res
}

// @lc code=end

func maxProfitGreedy(prices []int) int {
	pricesLen := len(prices)
	if pricesLen <= 1 {
		return 0
	}
	pre := prices[0]
	cur := prices[0]
	res := 0
	for i := 1; i < pricesLen; i++ {
		if prices[i] > cur {
			cur = prices[i]
		} else if prices[i] == cur {
			continue
		} else {
			res += cur - pre
			pre, cur = prices[i], prices[i]
		}
	}
	if pre != cur {
		res += cur - pre
	}
	return res
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfitGreedy([]int{1, 2, 3, 4, 5}))
}

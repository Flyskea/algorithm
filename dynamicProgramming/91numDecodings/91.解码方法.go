/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */
package main

import (
	"fmt"
	"strconv"
)

// @lc code=start
func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	if s[:1] == "0" {
		dp[1] = 0
	} else {
		dp[1] = 1
	}
	for i := 2; i <= len(s); i++ {
		lastNum, _ := strconv.Atoi(s[i-1 : i])
		if lastNum >= 1 && lastNum <= 9 {
			dp[i] += dp[i-1]
		}
		lastNum, _ = strconv.Atoi(s[i-2 : i])
		if lastNum >= 10 && lastNum <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}

// @lc code=end

func main() {
	fmt.Println(numDecodings("101002"))
}

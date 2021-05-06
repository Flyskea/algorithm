/*
 * @lc app=leetcode.cn id=583 lang=golang
 *
 * [583] 两个字符串的删除操作
 */
package main

import "fmt"

// @lc code=start

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDistance(word1 string, word2 string) int {
	word1Len, word2Len := len(word1), len(word2)
	if word1Len <= 0 || word2Len <= 0 {
		return 0
	}
	dp := make([][]int, word1Len+1)
	for i := range dp {
		dp[i] = make([]int, word2Len+1)
	}
	for i := 1; i <= word1Len; i++ {
		for j := 1; j <= word2Len; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	res := word2Len + word1Len - 2*dp[word1Len][word2Len]
	return res
}

// @lc code=end
// 不使用LCS的DP
// DP[i][j]表示前i和j个字符需要删除的数量
func minDistance1(word1 string, word2 string) int {
	word1Len, word2Len := len(word1), len(word2)
	if word1Len <= 0 || word2Len <= 0 {
		return 0
	}
	dp := make([][]int, word1Len+1)
	for i := range dp {
		dp[i] = make([]int, word2Len+1)
	}
	for i := 0; i <= word1Len; i++ {
		for j := 0; j <= word2Len; j++ {
			if i == 0 || j == 0 {
				// 如果任意字符串为0则删除所有字符串
				dp[i][j] = i + j
				// 完全匹配不用删除字符
			} else if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 不匹配 删除字符
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[word1Len][word2Len]
}
func main() {
	fmt.Println(minDistance1("sea", "eat"))
}

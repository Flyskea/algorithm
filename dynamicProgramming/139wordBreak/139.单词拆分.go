/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */
package main

import "fmt"

// @lc code=start
/* DFS + 记忆优化
func canBreak(start int, s string, wordMap map[string]bool, memo map[int]bool) bool {
	if start == len(s) {
		return true
	}
	for i := start + 1; i <= len(s); i++ {
		prefix := s[start:i]
		if wordMap[prefix] && canBreak(i, s, wordMap, memo) {
			memo[start] = true
			return true
		}
	}
	memo[start] = false
	return false
}

func wordBreak(s string, wordDict []string) bool {
	wordMap := map[string]bool{}
	for _, v := range wordDict {
		wordMap[v] = true
	}
	memo := make(map[int]bool)
	return canBreak(0, s, wordMap, memo)
} */

// DP
func wordBreak(s string, wordDict []string) bool {
	wordMap := map[string]bool{}
	sLen := len(s)
	for _, v := range wordDict {
		wordMap[v] = true
	}
	dp := make([]bool, sLen+1)
	dp[0] = true
	for i := 1; i <= sLen; i++ {
		for j := i - 1; j >= 0; j-- {
			if wordMap[s[j:i]] && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[sLen]
}

// @lc code=end

func main() {
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
}

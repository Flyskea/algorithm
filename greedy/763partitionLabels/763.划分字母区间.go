/*
 * @lc app=leetcode.cn id=763 lang=golang
 *
 * [763] 划分字母区间
 */
package main

import "fmt"

// @lc code=start
func partitionLabels(S string) []int {
	lastPos := make([]int, 26)
	var partition []int
	for i, v := range S {
		lastPos[v-'a'] = i
	}
	end, start := 0, 0
	for i, v := range S {
		if lastPos[v-'a'] > end {
			end = lastPos[v-'a']
		}
		if i == end {
			partition = append(partition, end-start+1)
			start = end + 1
		}
	}
	return partition
}

// @lc code=end

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}

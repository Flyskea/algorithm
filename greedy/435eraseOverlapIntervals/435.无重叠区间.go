/*
 * @lc app=leetcode.cn id=435 lang=golang
 *
 * [435] 无重叠区间
 */
package main

import (
	"fmt"
	"sort"
)

// @lc code=start
func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	var removed int
	prev := intervals[0][1]
	for i := 1; i < n; i++ {
		if intervals[i][0] < prev {
			removed++
		} else {
			prev = intervals[i][1]
		}
	}
	return removed
}

// @lc code=end

func main() {
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}}))
}

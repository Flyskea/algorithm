/*
 * @lc app=leetcode.cn id=452 lang=golang
 *
 * [452] 用最少数量的箭引爆气球
 */
package main

import (
	"fmt"
	"sort"
)

// @lc code=start
func findMinArrowShots(points [][]int) int {
	n := len(points)
	if n == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	prev, res := points[0][1], 1
	for i := 1; i < n; i++ {
		if points[i][0] > prev {
			prev = points[i][1]
			res++
		}
	}
	return res
}

// @lc code=end

func main() {
	fmt.Println(findMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}))
}

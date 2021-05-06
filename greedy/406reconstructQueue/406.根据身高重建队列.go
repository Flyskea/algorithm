/*
 * @lc app=leetcode.cn id=406 lang=golang
 *
 * [406] 根据身高重建队列
 */
package main

import (
	"fmt"
	"sort"
)

// @lc code=start
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return (people[i][0] > people[j][0]) || people[i][0] == people[j][0] && people[i][1] < people[j][1]
	})
	for i, v := range people {
		to := v[1]
		copy(people[to+1:i+1], people[to:i+1])
		people[to] = v
	}
	return people
}

// @lc code=end

func main() {
	fmt.Println(reconstructQueue([][]int{
		{7, 0}, {4, 4}, {7, 1},
		{5, 0}, {6, 1}, {5, 2},
	}))
}

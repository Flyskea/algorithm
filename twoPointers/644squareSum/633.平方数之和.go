/*
 * @lc app=leetcode.cn id=633 lang=golang
 *
 * [633] 平方数之和
 */
package main

import (
	"fmt"
	"math"
)

// @lc code=start
func judgeSquareSum(c int) bool {
	low, high := 0, int(math.Sqrt(float64(c)))
	for low <= high {
		sum := low*low + high*high
		if sum == c {
			return true
		} else if sum > c {
			high--
		} else {
			low++
		}
	}
	return false
}

// @lc code=end
func main() {
	fmt.Println(judgeSquareSum(4))
}

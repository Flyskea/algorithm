/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */
package main

import "fmt"

// @lc code=start
func mySqrt(x int) int {
	if x <= 0 {
		return x
	}
	low, high := 1, x
	var mid, res int
	for low <= high {
		mid = low + (high-low)/2
		res = x / mid
		if res == mid {
			return mid
		} else if res < mid {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return high
}

// @lc code=end
func main() {
	fmt.Println(mySqrt(1))
}

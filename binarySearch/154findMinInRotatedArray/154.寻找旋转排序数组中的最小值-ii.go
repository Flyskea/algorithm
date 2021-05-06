/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */
package main

import "fmt"

// @lc code=start
func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	var mid int
	for low <= high {
		mid = low + (high-low)/2
		if nums[mid] > nums[high] {
			low = mid + 1
		} else if nums[mid] < nums[high] {
			high = mid
		} else {
			high--
		}
	}
	return nums[low]
}

// @lc code=end

func main() {
	fmt.Println(findMin([]int{2, 2, 2, 0, 1}))
}

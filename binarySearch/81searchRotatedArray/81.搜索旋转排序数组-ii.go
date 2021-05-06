/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */
package main

import "fmt"

// @lc code=start
func search(nums []int, target int) bool {
	low, high := 0, len(nums)-1
	var mid int
	for low <= high {
		mid = low + (high-low)/2
		if target == nums[mid] {
			return true
		}
		if nums[mid] == nums[low] {
			low++
		} else if nums[mid] <= nums[high] {
			// 右区间是增序的
			if target > nums[mid] && target < nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			// 左区间是增序的
			if target > nums[low] && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return false
}

// @lc code=end

func main() {
	fmt.Println(search([]int{1, 0, 1, 1, 1}, 0))
}

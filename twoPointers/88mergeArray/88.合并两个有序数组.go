/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */
package main

import "fmt"

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2 := m-1, n-1
	p3 := m + n - 1
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p3] = nums1[p1]
			p1--
		} else {
			nums1[p3] = nums2[p2]
			p2--
		}
		p3--
	}
	for p1 >= 0 {
		nums1[p3] = nums1[p1]
		p1--
		p3--
	}
	for p2 >= 0 {
		nums1[p3] = nums2[p2]
		p2--
		p3--
	}
}

// @lc code=end

func main() {
	a := []int{0}
	merge(a, 0, []int{1}, 1)
	fmt.Println(a)
}

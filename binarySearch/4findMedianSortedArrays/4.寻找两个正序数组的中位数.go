/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */
package main

import "fmt"

// 根据中位数的定义，当 m+nm+n 是奇数时，
// 中位数是两个有序数组中的第 (m+n)/2(m+n)/2 个元素，
// 当 m+nm+n 是偶数时，中位数是两个有序数组中的第 (m+n)/2(m+n)/2 个元素和第 (m+n)/2+1(m+n)/2+1 个元素的平均值。

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
}

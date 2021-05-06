/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] 分发糖果
 */
package main

import "fmt"

/*
 * 两次遍历，先从左边遍历然后从右边遍历
func candy(ratings []int) int {
	ratingLen := len(ratings)
	if ratingLen < 2 {
		return ratingLen
	}
	nums := make([]int, ratingLen)
	for i := 0; i < ratingLen; i++ {
		nums[i] = 1
	}
	for i := 1; i < ratingLen; i++ {
		if ratings[i] > ratings[i-1] {
			nums[i] = nums[i-1] + 1
		}
	}
	for i := ratingLen - 1; i > 0; i-- {
		if ratings[i] < ratings[i-1] {
			if nums[i-1] < nums[i]+1 {
				nums[i-1] = nums[i] + 1
			}
		}
	}
	var res int
	for i := 0; i < ratingLen; i++ {
		res += nums[i]
	}
	return res
}
*/

// @lc code=start
// 常量遍历
// 如果当前同学比上一个同学评分高，说明我们就在最近的递增序列中，直接分配给该同学 \textit{pre} + 1pre+1 个糖果即可。
// 否则我们就在一个递减序列中，我们直接分配给当前同学一个糖果，并把该同学所在的递减序列中所有的同学都再多分配一个糖果，以保证糖果数量还是满足条件。
// 我们无需显式地额外分配糖果，只需要记录当前的递减序列长度，即可知道需要额外分配的糖果数量。
// 同时注意当当前的递减序列长度和上一个递增序列等长时，需要把最近的递增序列的最后一个同学也并进递减序列中。
func candy(ratings []int) int {
	n := len(ratings)
	ans, inc, dec, pre := 1, 1, 0, 1
	for i := 1; i < n; i++ {
		if ratings[i] >= ratings[i-1] {
			dec = 0
			if ratings[i] == ratings[i-1] {
				pre = 1
			} else {
				pre++
			}
			ans += pre
			inc = pre
		} else {
			dec++
			if dec == inc {
				dec++
			}
			ans += dec
			pre = 1
		}
	}
	return ans
}

// @lc code=end

func main() {
	fmt.Println(candy([]int{1, 2, 3, 4, 5}))
}

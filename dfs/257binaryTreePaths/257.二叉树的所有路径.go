/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
 */

package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	res := []string{}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}
	tmpLeft := binaryTreePaths(root.Left)
	for i := 0; i < len(tmpLeft); i++ {
		res = append(res, strconv.Itoa(root.Val)+"->"+tmpLeft[i])
	}
	tmpRight := binaryTreePaths(root.Right)
	for i := 0; i < len(tmpRight); i++ {
		res = append(res, strconv.Itoa(root.Val)+"->"+tmpRight[i])
	}
	return res
}

// @lc code=end

func main() {
	fmt.Println(binaryTreePaths(nil))
}

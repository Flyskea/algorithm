/*
 * @lc app=leetcode.cn id=310 lang=golang
 *
 * [310] 最小高度树
 */
package main

import "fmt"

// @lc code=start
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	if n == 2 {
		return []int{0, 1}
	}
	graph := make(map[int][]int, n)
	degree := make([]int, n)
	for i := range edges {
		u, v := edges[i][0], edges[i][1]
		graph[v] = append(graph[v], u)
		graph[u] = append(graph[u], v)
		degree[v]++
		degree[u]++
	}
	queue := []int{}
	for i := 0; i < n; i++ {
		if degree[i] == 1 {
			queue = append(queue, i)
		}
	}
	cnt := len(queue)
	for n > 2 {
		n -= cnt
		for cnt > 0 {
			tmp := queue[0]
			queue = queue[1:]
			degree[tmp] = 0
			for i := 0; i < len(graph[tmp]); i++ { //遍历当前节点的邻接节点
				if degree[graph[tmp][i]] != 0 {
					degree[graph[tmp][i]]--         //去掉与当前节点的关系
					if degree[graph[tmp][i]] == 1 { //如果度为1了 就加入队列
						queue = append(queue, graph[tmp][i])
					}
				}
			}
			cnt--
		}
		cnt = len(queue)
	}
	ans := []int{}
	ans = append(ans, queue...)
	return ans
}

// @lc code=end

func main() {
	fmt.Println(findMinHeightTrees(6, [][]int{
		{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4},
	}))
}

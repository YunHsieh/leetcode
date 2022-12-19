package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 1
	node := []*TreeNode{root}
	for {
		tempNode := []*TreeNode{root}
		for _, v := range node {
			if v.Right == nil && v.Left == nil {
				return depth
			}
			if v.Right != nil {
				tempNode = append(tempNode, v.Right)
			}
			if v.Left != nil {
				tempNode = append(tempNode, v.Left)
			}
		}
		depth++
		node = tempNode
	}
	return depth
}

func main() {

}

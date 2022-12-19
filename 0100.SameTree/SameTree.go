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
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    } else if p== nil || q==nil {
        return false
    }
    if p.Val == q.Val {
        return isSameTree(p.Right, q.Right) && isSameTree(p.Left, q.Left)
    } else {
        return false
    }
}

func main() {

}

package main

import (
	. "fmt"
	"testing"
)

// func parseNode(node *TreeNode) (TreeNode) {
//     if node.Left != nil {
//         Println(">>>left", node.Left.Val)
//         return parseNode(node.Left)
//     }
//     if node.Right != nil {
//         Println(">>>right", node.Right.Val)
//         return parseNode(node.Right)
//     }
//     return *node
// }

func TestListCovertToTreeNode(t *testing.T) {
    result := ListCovertToTreeNode([]int{10, 5, 3, 4, 8, 9, 1, 0, 1, 2})
    
    if result.Val != 10 {
        t.Error("fail")
    }
    if result.Left.Val != 5 {
        t.Error("fail")
    }
    if result.Left.Left.Val != 4 {
        t.Error("fail")
    }
    if result.Left.Left.Right.Val != 1 {
        t.Error("fail")
    }
}

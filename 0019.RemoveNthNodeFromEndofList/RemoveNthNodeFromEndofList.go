package main

import (
	. "leetcode/utils"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
    allNode := []*ListNode{head}
    node := head
    for node.Next != nil {
        node = node.Next
        allNode = append(allNode, node)
        if len(allNode) > n+1 {
            allNode = allNode[1:]
        }
    }
    if len(allNode) <= 1 {
        return nil
    } else if len(allNode) < n+1 {
        return allNode[1]
    }
    allNode[0].Next = allNode[0].Next.Next
    return head
}

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}
func nextNumber(currentNode *ListNode, list1 *ListNode, list2 *ListNode) int {
	newNode := &ListNode{}
	if list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			newNode = list2
			list2 = list2.Next
		} else {
			newNode = list1
			list1 = list1.Next
		}
		currentNode.Next = newNode
		currentNode = newNode
		return nextNumber(currentNode, list1, list2)
	}
	if list1 == nil {
		currentNode.Next = list2
	}
	if list2 == nil {
		currentNode.Next = list1
	}
	return 1
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    root := &ListNode{}
    nextNumber(root, list1, list2)
    return root.Next
}

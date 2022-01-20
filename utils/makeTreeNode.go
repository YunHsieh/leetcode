package utils

import (
	. "fmt"
)

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

type ListNode struct {
    Val int
    Next *ListNode
}

func ListCovertToTreeNode(numArray []int) TreeNode {
    if len(numArray) == 0 {
        return TreeNode{}
    }
    firstElement, numArray := numArray[0], numArray[1:]
    root := TreeNode{Val: firstElement}
    tempSlice := []*TreeNode{&root}
    currentNode, tempSlice := tempSlice[0], tempSlice[1:]
    isLeft := true
    for len(numArray) > 0 {
        firstElement, numArray = numArray[0], numArray[1:]
        if firstElement == 0 {
            if !isLeft {
                currentNode, tempSlice = tempSlice[0], tempSlice[1:]
            }
            isLeft = !isLeft
            continue
        }
        newNode := TreeNode{Val: firstElement}
        tempSlice = append(tempSlice, &newNode)

        if isLeft {
            isLeft = !isLeft
            currentNode.Left = &newNode
        } else {
            isLeft = !isLeft
            currentNode.Right = &newNode
            currentNode, tempSlice = tempSlice[0], tempSlice[1:]
        }
    }
    return root
}

func SliceToListNode(numArray []int) *ListNode {
    if len(numArray) == 0 {
        return nil
    }
    head := &ListNode{
        Val: numArray[0],
    }
    currentNode := head
    for i:=1; len(numArray)>i; i++ {
        nextNode := &ListNode{
            Val: numArray[i],
        }
        currentNode.Next, currentNode = nextNode, nextNode
    }
    return head
}

func CompareTwoListNode(node *ListNode,targetNode *ListNode) bool {
    for node != nil && targetNode != nil {
        if node.Val == targetNode.Val {
            node, targetNode = node.Next, targetNode.Next
        } else {
            return false
        }
    }
    if node != nil || targetNode != nil {
        return false
    }
    return true
}

func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    return &TreeNode{Val: nums[len(nums)/2], Left: sortedArrayToBST(nums[:len(nums)/2]), Right: sortedArrayToBST(nums[len(nums)/2+1:])}
}

func main() {
    result := ListCovertToTreeNode([]int{10, 5, 3, 4, 8, 9, 1, 0, 1, 2})
    Println(result.Val)
    Println(result.Left.Val)
    Println(result.Right.Val)
    Println(result.Left.Left.Val)
    Println(result.Left.Right.Val)
}

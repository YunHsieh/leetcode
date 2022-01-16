package main

import (
	. "fmt"
)

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
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

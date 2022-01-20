package utils

import (
	"testing"
)

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
    t.Log("> TestListCovertToTreeNode Success")
}

func TestSliceToListNode(t *testing.T) {
    nums := []int{10, 5, 3, 4, 8, 9, 1, 0, 1, 2}
    result := SliceToListNode(nums)
    
    for i:=0; len(nums)>i; i++ {
        if result.Val != nums[i] {
            t.Error("fail", result.Val, nums[i])
        }
        result = result.Next
    }
    t.Log("> TestSliceToListNode Success")
}

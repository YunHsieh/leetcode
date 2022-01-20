package main

import (
	utils "leetcode/utils"
	"testing"
)

func TestRemoveNthNodeFromEndofListScence1(t *testing.T) {
	node := utils.SliceToListNode([]int{2,3,1,4})
	result := RemoveNthFromEnd(node, 3)
	ans_node := utils.SliceToListNode([]int{2,1,4})
 
    if utils.CompareTwoListNode(result, ans_node) {
        t.Log("> TestRemoveNthNodeFromEndofListScence1 success")
    } else {
        t.Error("fail")
    }
}

func TestRemoveNthNodeFromEndofListScence2(t *testing.T) {
	node := utils.SliceToListNode([]int{1,2,3,4,5})
	result := RemoveNthFromEnd(node, 3)
	ans_node := utils.SliceToListNode([]int{1,2,4,5})
 
    if utils.CompareTwoListNode(result, ans_node) {
        t.Log("> TestRemoveNthNodeFromEndofListScence2 success")
    } else {
        t.Error("fail")
    }
}

func TestRemoveNthNodeFromEndofListScence3(t *testing.T) {
	node := utils.SliceToListNode([]int{1, 2})
	result := RemoveNthFromEnd(node, 2)
	ans_node := utils.SliceToListNode([]int{2})
 
    if utils.CompareTwoListNode(result, ans_node) {
        t.Log("> TestRemoveNthNodeFromEndofListScence3 success")
    } else {
        t.Error("fail")
    }
}

func TestRemoveNthNodeFromEndofListScence4(t *testing.T) {
	node := utils.SliceToListNode([]int{1})
	result := RemoveNthFromEnd(node, 1)
	ans_node := utils.SliceToListNode([]int{})
 
    if utils.CompareTwoListNode(result, ans_node) {
        t.Log("> TestRemoveNthNodeFromEndofListScence4 success")
    } else {
        t.Error("fail")
    }
}

func TestRemoveNthNodeFromEndofListScence5(t *testing.T) {
	node := utils.SliceToListNode([]int{1, 2})
	result := RemoveNthFromEnd(node, 1)
	ans_node := utils.SliceToListNode([]int{1})
 
    if utils.CompareTwoListNode(result, ans_node) {
        t.Log("> TestRemoveNthNodeFromEndofListScence5 success")
    } else {
        t.Error("fail")
    }
}

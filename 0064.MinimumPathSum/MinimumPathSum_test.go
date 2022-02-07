package main

import (
	"testing"
)

func TestMinimumPathSumScence1(t *testing.T) {
	param1 := [][]int{}
	param1 = append(param1, []int{1,3,1})
	param1 = append(param1, []int{1,5,1})
	param1 = append(param1, []int{4,2,1})
    result := MinPathSum(param1)
	ans := 7
    if result == ans {
        t.Log("success")
    } else {
        t.Error("fail", result)
    }
}


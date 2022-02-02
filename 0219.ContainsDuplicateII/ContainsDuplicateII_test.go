package main

import "testing"

func TestContainsDuplicateIIScence1(t *testing.T) {
    result := ContainsNearbyDuplicate([]int{1,2,3,1,2,3}, 2)
	ans := false
    if result == ans {
        t.Log("success")
    } else {
        t.Error("fail", result)
    }
}

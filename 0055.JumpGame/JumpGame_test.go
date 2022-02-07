package main

import (
	"testing"
)

func TestCanJumpScence1(t *testing.T) {
    result := CanJump([]int{2,3,1,1,4})
    ans := true

    if result == ans {
        t.Log("success")
    } else {
        t.Error("fail: ", result)
    }
}

func TestCanJumpScence2(t *testing.T) {
    result := CanJump([]int{1})
    ans := true

    if result == ans {
        t.Log("success")
    } else {
        t.Error("fail: ", result)
    }
}

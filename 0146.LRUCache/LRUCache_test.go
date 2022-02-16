package main

import (
	"testing"
)

func TestLRUCacheScence1(t *testing.T) {
    result := testFunc()
    ans := 1

    if result == ans {
        t.Log("success")
    } else {
        t.Error("fail: ", result)
    }
}

func testFunc() int {
    return 1
}
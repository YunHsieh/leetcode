package main

import (
	"testing"
)

func TestLongestCommonSubsequenceScene1(t *testing.T) {
    result := LongestCommonSubsequence("oxcpqrsvwf", "shmtulqrypy")

    if result == 2 {
        t.Log("success")
    } else {
        t.Error("fail")
    }
}

func TestLongestCommonSubsequenceScene2(t *testing.T) {
    result := LongestCommonSubsequence("abcde", "abc")

    if result == 3 {
        t.Log("success")
    } else {
        t.Error("fail")
    }
}

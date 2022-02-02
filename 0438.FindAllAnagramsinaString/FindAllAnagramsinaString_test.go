package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindAnagramsScence1(t *testing.T) {
    result := FindAnagrams("cbaebabacd", "abc")
	ans := []int{0,6}
	sort.Ints(result)
	sort.Ints(ans)
    if reflect.DeepEqual(result, ans) {
        t.Log("success")
    } else {
        t.Error("fail", result)
    }
}

func TestFindAnagramsScence2(t *testing.T) {
    result := FindAnagrams("ababababab", "aab")
	ans := []int{0,2,4,6}
	sort.Ints(result)
	sort.Ints(ans)
    if reflect.DeepEqual(result, ans) {
        t.Log("success")
    } else {
        t.Error("fail", result)
    }
}

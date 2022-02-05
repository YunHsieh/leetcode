package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestTopKFrequentScence1(t *testing.T) {
    result := TopKFrequent([]int{1,1,1,2,2,3333}, 2)
    ans := []int{1,2}
    sort.Ints(ans)
    sort.Ints(result)
    if reflect.DeepEqual(result, ans) {
        t.Log("success")
    } else {
        t.Error("fail", result, ans)
    }
}

func TestTopKFrequentScence2(t *testing.T) {
    result := TopKFrequent([]int{6,0,1,4,9,7,-3,1,-4,-8,4,-7,-3,3,2,-3,9,5,-4,0}, 6)
    ans := []int{-3,-4,0,1,4,9}
    sort.Ints(ans)
    sort.Ints(result)
    fmt.Printf("%v %v", result, ans)
    if reflect.DeepEqual(result, ans) {
        t.Log("success")
    } else {
        t.Error("fail", result, ans)
    }
}

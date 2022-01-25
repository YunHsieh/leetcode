/*
https://leetcode.com/problems/top-k-frequent-elements/discuss/353045/Golang-16ms-solution-using-a-heap
*/

package main

import (
	. "container/heap"
)

type ElementStruct struct {
    Value int
    Frequent int
}

type MinHeap []*ElementStruct

func (h MinHeap) Less(i, j int) bool {
    return h[i].Frequent > h[j].Frequent
}

func (h MinHeap) Len() int {
    return len(h)
}

func (h MinHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(*ElementStruct))
}

func (h *MinHeap) Pop() interface{} {
    tmp := *h
    n := len(tmp)
    *h = tmp[:n-1]
    return tmp[n-1]
}

func TopKFrequent(nums []int, k int) []int {
    if len(nums) == 0 {
        return nums
    }
    counter := map[int]int{}
    for i:=0; len(nums)>i; i++ {
        counter[nums[i]] += 1
    }
    frequentHeap := MinHeap{}
    for Key, Val := range counter {
        newElement := ElementStruct{
            Value: Key,
            Frequent: Val,
        }
        frequentHeap = append(frequentHeap, &newElement)
    }

    Init(&frequentHeap)
    result := []int{}
    for i:=0; i<k; i++ {
        result = append(result, frequentHeap[0].Value)
        Pop(&frequentHeap)
    }
    return result
}

package main

import (
	"container/heap"
	"fmt"
)

type myHeap []int

func (h myHeap) Len() int{
    return len(h)
}

func (h myHeap) Swap(i, j int){
    h[i], h[j] = h[j], h[i]
}

func (h myHeap) Less(i, j int) bool{
    // heap 預設是實作 min heap
    // 因為我們要實作 max heap
    // 所以把 less 的答案倒轉
    return h[i] > h[j]
}

// Push 要實作的是就是
// 當 heap 新增元素要把新元素加入哪裡 (即 *h[Len(*h)])
// 不需要實作 swim
func (h *myHeap) Push(x interface{}){
    *h = append(*h, x.(int))
}

// Pop 要實作的是
// 當有元素被移除時，要用哪一個元素去替補
// 通常都是拿最後一個 (即 *h[Len(*h)-1])
// 不需要實作 sink
func (h *myHeap) Pop() interface{}{
    fmt.Printf("%v", h)
    old := *h
    tmp := old[len(*h)-1]
    // 切片重切，因為 heap 套件在內部實作時是靠
    // 切片的長度在定位的，跟我們先前的方式不太一樣
    // 所以一定要重切
    *h = old[0: len(*h)-1]
    fmt.Printf("%v", h)
    return tmp
}

func main(){
    h := &myHeap{-5, -1, 1, 2, 3, 4, 8}
    // init 會將初始的值先做成堆
    // 這樣就不用像我們之前實作時得從空陣列(切片)開始加了
    heap.Init(h)
    heap.Push(h, 9)
    fmt.Printf("%v\n", h)
    fmt.Printf("%d\n", heap.Pop(h))
    fmt.Printf("%d\n", heap.Pop(h))
    fmt.Printf("%d\n", heap.Pop(h))
}
package main

type arrData struct {
    Val int
}

type Minheap []*arrData

func (a Minheap) Init() {
    for i:=len(a)/2-1; i >= 0; i-- {
        a.down(i)
    }
}

func (a Minheap) Len() int {
	return len(a)
}

func (a Minheap) Less(i, j int) bool {
    return a[i].Val < a[j].Val
}

func (a *Minheap) Pop() int {
	old := *a
	n := old.Len()
	a.Swap(0, n-1)
	*a = old[:n-1]
	a.down(0)
	return old[0].Val
}

func (a Minheap) down(i int) {
    n := len(a)
    for {
        j1 := i*2+1
        if j1 >= n || j1 < 0 {
            break
        }
        j := j1
        if j2 := j1+1; j2 < n && a.Less(j1, j2) {
            j = j2
        }
        if !a.Less(i, j) {
            break
        }
        a.Swap(i, j)
        i = j
    }
    
}

func (a Minheap) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}

func FindKthLargest(nums []int, k int) int {
	h := Minheap{}
    for _, val := range nums {
        h = append(h, &arrData{Val:val})
    }
	h.Init()
	
	result := h[0].Val
	for i:=0; i < k-1; i++ {
		result = h.Pop()
	}
	return result
}

// another answer
// func findKthLargest(nums []int, k int) int {
//     sort.Ints(nums)
//     return nums[len(nums)-k]
// }

func main() {
	FindKthLargest([]int{3,2,1,5,6,4}, 2)
	FindKthLargest([]int{1}, 1)
	FindKthLargest([]int{2,1}, 2)
}
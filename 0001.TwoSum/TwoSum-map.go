package main

func twoSum(nums []int, target int) []int {
    mapping := map[int]int{}
    for i, x := range nums {
        if val, ok := mapping[target-x]; ok{
            return []int{i, val}
        }
        mapping[x] = i
    }
        return []int{}
}

func main () {
    a := twoSum([]int{1,3,4,5}, 4)
    println(a)
}

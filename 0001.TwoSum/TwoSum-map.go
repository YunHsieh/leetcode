package main

func twoSumMap(nums []int, target int) []int {
    mapping := map[int]int{}
    for i, x := range nums {
        if val, ok := mapping[target-x]; ok{
            return []int{i, val}
        }
        mapping[x] = i
    }
        return []int{}
}

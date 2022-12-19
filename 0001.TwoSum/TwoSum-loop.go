package main

func twoSumLoop(nums []int, target int) []int {
    for i:=0; i < len(nums); i++ {
        ans := target - nums[i]
        for j:=i+1; j < len(nums); j++ {
            if ans == nums[j] {
                return []int{i, j}
            }
        }
    }
    return []int{}
}

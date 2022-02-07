package main

func CanJump(nums []int) bool {
    n := len(nums) - 1
    for i:=n; i>=0; i-- {
        if nums[i]+i >= n {
            n = i
        }
    }
    return n == 0
}

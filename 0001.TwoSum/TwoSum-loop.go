package main

import (
    . "fmt"
) 

func twoSum(nums []int, target int) []int {
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

func main () {
    a := twoSum([]int{1,3,4,5}, 4)
    Println(a)

}

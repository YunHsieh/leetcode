package main

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
    result := [][]int{}
    sort.Ints(nums)
    for i:=0; len(nums)-2>i; i++ {
        if i!=0 && nums[i-1] == nums[i] {
			continue
		}
        var start, end int = i+1, len(nums)-1
        for ;start < end; {
            ans := nums[i] + nums[start] + nums[end]
            if ans > 0 {
                end--
            } else if ans < 0 {
                start++
            } else if ans == 0 {
                newElement := []int{nums[i], nums[start], nums[end]}
                nextElement := start 
                result = append(result, newElement)
                for nextElement < end && nums[start] == nums[nextElement] {
                    nextElement++
                }
                start = nextElement
            }
        }
    }
    return result
}

func main() {

}

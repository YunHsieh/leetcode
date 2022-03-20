/*
Rob detail explain:
https://leetcode.com/problems/house-robber/discuss/156523/From-good-to-great.-How-to-approach-most-of-DP-problems
*/
package HouseRobber

func max(m, n int) int {
    if m > n {
        return m
    }
    return n
}

func rob(nums []int) int {
    n := len(nums)
    result := make([]int, n+1)
    if n == 0 {
        return 0
    }
    if n == 1 {
        return nums[0]
    }
    if n == 2 {
        return max(nums[0], nums[1])
    }
    result[0] = nums[0]
    result[1] = max(nums[0],nums[1])
    
    for i:=2; i<n; i++ {
        result[i] = max(result[i-2]+nums[i], result[i-1])
    }
    return max(result[n-2], result[n-1])
}

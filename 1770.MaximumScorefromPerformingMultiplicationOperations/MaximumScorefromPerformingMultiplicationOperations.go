package main

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func dfs(dp [][]int, nums []int, multipliers []int, l, r int) int {
        if r >= len(multipliers) {
            return 0
        }
        if dp[l][r] == 0 {
            dp[l][r] = max(nums[l]*multipliers[r] +dfs(dp, nums, multipliers, l+1, r+1), nums[len(nums)-1-(r-l)]*multipliers[r]+dfs(dp, nums, multipliers, l, r+1))
        }
        return dp[l][r]
    }

func maximumScore(nums []int, multipliers []int) int {
    dp := make([][]int, len(multipliers))
    for i := range dp {
        dp[i] = make([]int, len(multipliers))
    }
    return dfs(dp, nums, multipliers, 0, 0)
}

func main() {

}

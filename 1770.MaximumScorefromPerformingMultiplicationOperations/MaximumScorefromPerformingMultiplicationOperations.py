class Solution:
    def maximumScore(self, nums: List[int], multipliers: List[int]) -> int:
        dp = [[0 for _ in multipliers] for _ in multipliers]
        m, n = len(multipliers), len(nums)
        def dfs(l, r):
            if r >= m: return 0
            if not dp[l][r]:
                dp[l][r] = max(nums[l]*multipliers[r] + dfs(l+1, r+1), 
                               nums[n-1-(r-l)]*multipliers[r] + dfs(l, r+1))
            return dp[l][r]

        return dfs(0, 0)

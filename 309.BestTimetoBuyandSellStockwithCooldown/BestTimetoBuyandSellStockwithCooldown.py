class Solution:
    def maxProfit(self, prices: list[int]) -> int:
        n = len(prices)
        if n < 2: 
            return 0
        dp = [0 for _ in range(n)]
        for i in range(1, n):
            dp[i] = dp[i-1]
            for j in range(i):
                tmp = prices[i] - prices[j]
                tmp += dp[j-2] if j > 1 else 0
                dp[i] = max(dp[i], tmp)
        return dp[n-1]

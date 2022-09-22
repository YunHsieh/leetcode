class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        results = [float("inf")]*(amount+1)
        results[0] = 0
        for i in range(1, amount+1):
            for coin in coins:
                if coin <= i:
                    results[i] = min(results[i], results[i-coin] + 1)
        return results[amount] if results[amount] != float("inf") else -1

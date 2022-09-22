class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        coins = sorted(coins, reverse=True)
        results = [-1]
        def find_amount(data, amount):
            if amount == 0:
                results[0] = data
            if amount <= 0:
                return amount
            for coin in coins:
                if results[0] != -1:
                    return results[0]
                if find_amount(data+1, amount-coin) == 0:
                    return data+1
        find_amount(0, amount)
        return results[0]

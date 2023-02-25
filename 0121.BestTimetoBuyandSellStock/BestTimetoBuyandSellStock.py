'''
Time complexity: O(N)
Space complexity: O(1)

Runtime Faster than 90%+
'''


class Solution:
    def maxProfit(self, prices: list[int]) -> int:
        min_p, result = prices[0], 0
        for price in prices:
            if price <= min_p:
                min_p = price
            else:
                result = max(price-min_p, result)
        return result

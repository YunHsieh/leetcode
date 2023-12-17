'''
Time complexity: O(N)
Space complexity: O(N)
Runtime: 69%+ 
Memory: 79%+
'''

from typing import List


class Solution:
    def minCostClimbingStairs(self, cost: List[int]) -> int:
        if len(cost) <= 2:
            return min(cost[0], cost[1])
        result = [cost[0], cost[1]]
        for i, price in enumerate(cost[2:], start=2):
            result.append(min(result[i-1], result[i-2]) + price)
        return min(result[-1], result[-2])

'''
Time complexity: O(N)
Space complexity: O(N)
Runtime: 69%+ 
Memory: 79%+
'''

from typing import List


class Solution:
    def minCostClimbingStairs(self, cost: List[int]) -> int:
        prev, curr = cost[0], cost[1]
        for price in cost[2:]:
            prev, curr = curr, min(prev, curr) + price
        return min(prev, curr)

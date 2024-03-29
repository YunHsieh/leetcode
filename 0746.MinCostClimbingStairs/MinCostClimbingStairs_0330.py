from typing import List


class Solution:
    def minCostClimbingStairs(self, cost: List[int]) -> int:
        result = [cost[0], cost[1]]
        for i, num in enumerate(cost[2:]):
            result.append(min(result[i]+num, result[i+1]+num))
        return min(result[-1], result[-2])

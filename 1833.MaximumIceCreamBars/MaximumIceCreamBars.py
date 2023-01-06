class Solution:
    def maxIceCream(self, costs: list[int], coins: int) -> int:
        costs.sort()
        result = 0
        for cost in costs:
            coins -= cost
            if coins < 0:
                return result
            result += 1
        return result

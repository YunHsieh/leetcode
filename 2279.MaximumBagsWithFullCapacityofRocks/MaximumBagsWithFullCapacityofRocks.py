class Solution:
    def maximumBags(self, capacity: list[int], rocks: list[int], additionalRocks: int) -> int:
        result = 0
        remains = []
        for i in range(len(capacity)):
            if (tmp := capacity[i] - rocks[i]) == 0:
                result += 1
            else:
                remains.append(tmp)
        remains.sort()
        total = 0
        while remains:
            total += remains.pop(0)
            if total <= additionalRocks:
                result += 1
        return result

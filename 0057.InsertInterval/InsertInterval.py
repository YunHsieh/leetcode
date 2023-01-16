class Solution:
    def insert(self, intervals: list[list[int]], newInterval: list[int]) -> list[list[int]]:
        intervals = intervals
        start, end = newInterval
        tmp = [start, end]
        result = []
        for a, b in intervals:
            if a > end or b < start:
                result.append([a, b])
            else:
                tmp.extend([a, b])
        result.append([min(tmp), max(tmp)])
        return sorted(result)

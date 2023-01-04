class Solution:
    def minimumRounds(self, tasks: list[int]) -> int:
        level_count = {}
        result = 0
        for level in tasks:
            level_count[level] = level_count.get(level, 0) + 1
        for k, v in level_count.items():
            if v < 2:
                return -1
            result += v // 3
            if v % 3 != 0:
                result += 1
        return result

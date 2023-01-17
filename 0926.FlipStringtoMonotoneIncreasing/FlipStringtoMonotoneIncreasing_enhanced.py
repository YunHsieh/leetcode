class Solution:
    def minFlipsMonoIncr(self, s: str) -> int:
        digit_count = sum(1 for i in s if i == '0')
        result = digit_count
        if not result: return result
        ones = 0
        for i in s:
            if i == '0':
                digit_count -= 1
            else:
                ones += 1
            result = min(result, digit_count+ones)
        return result

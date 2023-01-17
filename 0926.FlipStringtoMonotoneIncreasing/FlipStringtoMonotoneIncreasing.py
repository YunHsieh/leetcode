class Solution:
    def minFlipsMonoIncr(self, s: str) -> int:
        result = float("inf")
        n = len(s)
        digit_0  = [0 for i in range(n)]
        digit_1 = [0 for i in range(n)]
        for i in range(n):
            if i != 0:
                digit_0[i], digit_1[n-i-1] = digit_0[i-1], digit_1[n-i]
            if s[i] == '1':
                digit_0[i] += 1
            if s[n-i-1] == '0':
                digit_1[n-i-1] += 1
        for i in range(n-1):
            result = min(result, digit_0[i] + digit_1[i+1])
        return min(result, digit_0[-1], digit_1[0])

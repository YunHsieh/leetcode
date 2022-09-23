class Solution:
    def concatenatedBinary(self, n: int) -> int:
        result, m = 0, 10**9 + 7
        for i in range(1, n+1):
            result = (result * (1 << (len(bin(i)) - 2)) + i) % m
        return result

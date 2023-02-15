'''
Time complexity: O(N)
Space complexity: O(N)
Runtime Faster than 95%+
'''

class Solution:
    def addToArrayForm(self, num: list[int], k: int) -> list[int]:
        result = []
        for i, data in enumerate(num[::-1]):
            k += data
            result.append(k - (k//10*10))
            k //= 10
        while k >= 1:
            result.append(k - (k//10*10))
            k //= 10
        return result[::-1]

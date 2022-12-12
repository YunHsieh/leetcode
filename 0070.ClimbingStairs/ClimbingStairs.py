'''
This is a DP question.
But it can be solved by Fibonacci algorithm.
'''

class Solution:
    def climbStairs(self, n: int) -> int:
        if n <= 3:
            return n
        results = [2, 3]
        for i in range(3, n, 1):
            results[0], results[1] = results[1], sum(results)
        return results[1]

'''
Time complexity: O(N^2)
Space complexity: O(N)

Time Limited Exceed
'''

class Solution:
    def jump(self, nums: list[int]) -> int:
        n = len(nums)
        if n <= 1:
            return 0
        result = []
        def dfs(s, e, t):
            if e >= n:
                result.append(t)
                return
            for i, num in enumerate(nums[s:e]):
                dfs(s+i+1, num+i+s+1, t+1)
        dfs(1, 1+nums[0], 1)
        return min(result or [0])

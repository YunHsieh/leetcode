'''
Time complexity: O(N^2)
Space complexity: O(1)
Runtime Faster than 92%+
'''

class Solution:
    def jump(self, nums: list[int]) -> int:
        n = len(nums)
        if n <= 1:
            return 0
        result = 1
        s, e = 0, 1
        while e <= n:
            for i in range(s, e):
                if i+nums[i]+1 >= n:
                    return result
                if i+nums[i] >= e:
                    e = i+nums[i]+1
                    s = i+1
            result += 1
        return result

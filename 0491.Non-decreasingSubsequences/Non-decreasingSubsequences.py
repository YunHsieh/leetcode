'''
NOTE: if the solution use the list as the storage.
It'll go over the limited.
'''

class Solution:
    def findSubsequences(self, nums: list[int]) -> list[list[int]]:
        result = set()
        def increase(i, data):
            if len(data) > 1:
                result.add(tuple(data))
            if i >= len(nums):
                return
            if not data or nums[i] >= data[-1]:
                increase(i+1, [*data, nums[i]])
            increase(i+1, data)
        increase(0, [])
        return result

from typing import List


class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        result = []
        def dfs(mynums, data):
            if len(data) > len(nums):
                return
            result.append(data)
            for i in range(len(mynums)):
                dfs(mynums[i+1:], [mynums[i]]+data)
        dfs(nums, [])
        return result

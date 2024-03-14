from typing import List


class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        result = []
        def dfs(mynums, tmp):
            if len(tmp) == len(nums):
                result.append(tmp)
                return
            for i, num in enumerate(mynums):
                dfs(mynums[:i]+mynums[i+1:], tmp+[num])
        dfs(nums, [])
        return result

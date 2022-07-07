from typing import List


class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        results = []
        candidates = sorted(candidates)
        def backtrack(data: List[int]=[], total: int=0, index: int=0):
            if total > target:
                return

            if total == target:
                results.append(data.copy())
                return 

            for i in range(index, len(candidates)):
                data.append(candidates[i])
                backtrack(data, total+candidates[i], i)
                data.pop()
        backtrack()
        return results

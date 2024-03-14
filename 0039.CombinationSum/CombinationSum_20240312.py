from typing import List


class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        result = []
        def backtrack(mycandidates, tmp, count):
            for i, data in enumerate(mycandidates):
                summary = count+data
                if summary == target:
                    result.append(tmp+[data])
                if summary < target:
                    backtrack(mycandidates[i:], tmp+[data], count+data)

        backtrack(candidates, [], count=0)
        return result
from typing import List


class Solution:
    def partition(self, s: str) -> List[List[str]]:
        result = []
        n = len(s)
        def backtrack(mys, tmp, index=0):
            if index == n:
                result.append(tmp)
                return
            for i in range(1, len(mys)+1):
                if mys[:i] == mys[:i][::-1]:
                    backtrack(mys[i:], tmp+[mys[:i]], index+i)
        backtrack(s, [])
        return result

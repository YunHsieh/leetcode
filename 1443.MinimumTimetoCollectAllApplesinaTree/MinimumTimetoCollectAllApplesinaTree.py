"""
I posted the solution on the leetcode for the first times.
https://leetcode.com/problems/minimum-time-to-collect-all-apples-in-a-tree/solutions/3034712/python3-dfs-runtime-99-36-memory-94-87/
"""


class Solution:
    def minTime(self, n: int, edges: List[List[int]], hasApple: List[bool]) -> int:
        def dfs(route):
            if not route:
                return
            result.add(route)
            dfs(road.get(route))
            return route
        road, result = {}, set()
        for i, k in edges:
            if k in road:
                road[i] = k
            else:
                road[k] = i
        [dfs(i) for i, apple in enumerate(hasApple) if apple]
        return len(result)*2

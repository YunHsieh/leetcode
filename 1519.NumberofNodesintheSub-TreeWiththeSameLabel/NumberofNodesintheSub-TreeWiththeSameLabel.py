from collections import defaultdict


class Solution:
    def countSubTrees(self, n: int, edges: list[list[int]], labels: str) -> list[int]:
        graph, label_ct, ans = defaultdict(list), defaultdict(int), [0] * n
        for a, b in edges:
            graph[a].append(b)
            graph[b].append(a)
        
        def dfs(node=0, par=None):
            prev = label_ct[labels[node]]
            label_ct[labels[node]] += 1
            for child in graph[node]:
                if child != par: 
                    dfs(child, node)
            ans[node] = label_ct[labels[node]] - prev
        dfs()
        return ans

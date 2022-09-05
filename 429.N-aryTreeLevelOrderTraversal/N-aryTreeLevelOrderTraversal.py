"""
# Definition for a Node.
class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children
"""

class Solution:
    def levelOrder(self, root: 'Node') -> List[List[int]]:
        def bfs(node, layer):
            if not node:
                return
            if result.get(layer):
                result[layer].append(node.val)
            else:
                result[layer] = [node.val]
            for i in node.children:
                bfs(i, layer+1)
        result = {}
        bfs(root, 0)
        return result.values()

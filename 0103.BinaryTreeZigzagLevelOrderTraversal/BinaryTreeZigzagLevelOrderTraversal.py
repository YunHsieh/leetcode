'''
Time complexity: O(N)
Space complexity: O(N)

Runtime Faster than 76%+
'''

# Definition for a binary tree node.
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def zigzagLevelOrder(self, root: Optional[TreeNode]) -> list[list[int]]:
        data = {}
        def traversal(node, level=0):
            if not node:
                return
            if level % 2 == 0:
                data[level] = data.get(level,[]) + [node.val]
            else:
                data[level] = [node.val] + data.get(level,[])
            traversal(node.left, level+1)
            traversal(node.right, level+1)
        traversal(root)        
        return data.values()

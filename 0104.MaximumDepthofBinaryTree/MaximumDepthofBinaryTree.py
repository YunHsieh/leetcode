'''
Time complexity: O(1)
Space complexity: O(1)

Runtime Faster than 96%+
'''
# Definition for a binary tree node.
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        def dfs(node, depth=0):
            if not node:
                return depth
            return max(dfs(node.left, depth+1), dfs(node.right, depth+1))
        return dfs(root)

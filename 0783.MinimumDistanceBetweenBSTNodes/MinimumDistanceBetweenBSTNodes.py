'''
Time complexity: O(N)
Space complexity: O(N)

Runtime Faster than 95%+
'''

from typing import Optional


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def minDiffInBST(self, root: Optional[TreeNode]) -> int:
        data = []
        result = float("inf")
        def dfs(node):
            if not node:
                return
            dfs(node.left)
            data.append(node.val)
            dfs(node.right)
        dfs(root)
        for i in range(1, len(data)):
            result = min(result, data[i] - data[i-1])
        return result

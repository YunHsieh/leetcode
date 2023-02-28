# Definition for a binary tree node.
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def findDuplicateSubtrees(self, root: Optional[TreeNode]) -> list[Optional[TreeNode]]:
        subtrees, result = {}, []
        def dfs(node):
            if not node: return
            l, r = dfs(node.left), dfs(node.right)
            subtree_id = f'l{l}{node.val}{r}r'
            subtrees[subtree_id] = subtrees.get(subtree_id, 0) + 1
            if subtrees[subtree_id] == 2:
                result.append(node)
            return subtree_id
        dfs(root)
        return result

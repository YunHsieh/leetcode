# Definition for a binary tree node.
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def rightSideView(self, root: Optional[TreeNode]) -> List[int]:
        if not root:
            return []
        right_element = {}
        def dfs(node: Optional[TreeNode], level):
            if not node:
                return
            if right_element.get(level[0], None) is not None:
                if right_element[level[0]][1] < level[1]:
                    right_element[level[0]] = [node.val, level[1]]
            else:
                right_element[level[0]] = [node.val, level[1]]
            dfs(node.right, [level[0]+1, level[1]*2])
            dfs(node.left, [level[0]+1, level[1]*2-1])
        dfs(root, [0, 1])
        return [i[0] for i in right_element.values()]

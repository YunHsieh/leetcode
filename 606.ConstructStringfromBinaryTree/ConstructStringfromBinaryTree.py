# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def tree2str(self, root: Optional[TreeNode]) -> str:
        def dfs(node):
            res = ''
            if not node:
                return res
            left = dfs(node.left)
            right = dfs(node.right)
            if left or right:
                res += f"({left})"
            if right:
                res += f"({right})"
            return f'{node.val}' + res
        return dfs(root)

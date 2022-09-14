# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def pseudoPalindromicPaths (self, root: Optional[TreeNode]) -> int:
        def dfs(node, data):
            if not node:
                return 0
            if node.val in data:
                data.remove(node.val)
            else:
                data.add(node.val)
            if node.left is None and node.right is None and len(data) <= 1:
                    return 1
            return dfs(node.left, data.copy()) + dfs(node.right, data.copy())

        return dfs(root, set())

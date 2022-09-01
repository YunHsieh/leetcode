# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def dfs(self, node, current_max):
        if not node:
            return
        if node.val >= current_max:
            self.count+=1
            current_max = node.val
        self.dfs(node.left, current_max)
        self.dfs(node.right, current_max)

    def goodNodes(self, root: TreeNode) -> int:
        self.count = 0
        self.dfs(root, root.val)
        return self.count

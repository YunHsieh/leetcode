# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def minDiffInBST(self, root: Optional[TreeNode]) -> int:
        def traverse(node: TreeNode, low: int, high: int) -> int:
            if not node:
                return high - low
            left = traverse(node.left, low, node.val)
            right = traverse(node.right, node.val, high)
            return min(left, right)

        return traverse(root, -1 << 31, (1 << 31) - 1)

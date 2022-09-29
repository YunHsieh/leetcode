# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def pathSum(self, root: Optional[TreeNode], targetSum: int) -> List[List[int]]:
        results = []
        def dfs(node, data):
            if not node:
                return 
            dfs(node.left, data + [node.val])
            dfs(node.right, data + [node.val])
            if not node.left and not node.right and sum(data) + node.val == targetSum:
                results.append(data + [node.val])
        dfs(root ,[])
        return results

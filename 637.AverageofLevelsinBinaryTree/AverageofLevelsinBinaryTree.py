# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def averageOfLevels(self, root: Optional[TreeNode]) -> List[float]:
        def dfs(node, layer):
            if not node:
                return
            result[layer] = result.get(layer, []) + [node.val]
            dfs(node.right, layer+1)
            dfs(node.left, layer+1)

        result = {}
        dfs(root, 0)
        return [sum(i)/len(i) for i in result.values()]

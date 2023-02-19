'''
refer: https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/solutions/3204034/python3-26-ms-faster-than-97-08-of-python3/?page=2
'''
from collections import deque
from typing import Optional


class Solution:
    def zigzagLevelOrder(self, root: Optional[TreeNode]) -> list[list[int]]:
        if not root:
            return []

        queue = deque([root])
        result, direction = [], 1

        while queue:
            level = [node.val for node in queue]
            if direction == -1:
                level.reverse()
            result.append(level)
            direction *= -1

            for i in range(len(queue)):
                node = queue.popleft()
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)

        return result

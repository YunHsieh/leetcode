"""
Runtime: 141ms, 99.96%
Memory Usage: 13.9MB, 94.69%

Referance: https://leetcode.com/problems/maximum-length-of-repeated-subarray/discuss/2599253/Python3-Runtime%3A-178-ms-faster-than-99.92-or-Memory%3A-13.8-MB-less-than-99.81
"""
import os

class Solution:
    def findDuplicate(self, paths: List[str]) -> List[List[str]]:
        results = {}
        for i in paths:
            path, flies = i.split(' ', 1)
            for file in flies.split(' '):
                name, content = file.split('(', 1)
                if not results.get(content):
                    results[content] = [os.path.join(path, name)]
                else:
                    results[content].append(os.path.join(path, name))
        return [i for i in results.values() if len(i) > 1]

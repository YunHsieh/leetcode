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

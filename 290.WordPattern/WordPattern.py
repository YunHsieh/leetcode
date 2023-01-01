"""
T: 95.33%
M: 99.8%
"""

class Solution:
    def wordPattern(self, pattern: str, s: str) -> bool:
        s = s.split(" ")
        if len(pattern) != len(s) or len(set(s)) != len(set(pattern)):
            return False
        mapping = {}
        for i in range(len(s)):
            mapping[pattern[i]] = mapping.get(pattern[i], s[i])
            if mapping[pattern[i]] != s[i]:
                return False
        return True

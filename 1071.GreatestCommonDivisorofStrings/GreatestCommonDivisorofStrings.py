'''
This solution is try to find the all divitors.
'''

class Solution:
    def gcdOfStrings(self, str1: str, str2: str) -> str:
        n, m = len(str1), len(str2)
        text, min_len = min(str1, str2), min(n, m)
        for i in range(min_len, 0, -1):
            if m % i == 0 and n % i == 0:
                text = text[:i]
                if text * (n//len(text)) == str1 and text * (m//len(text)) == str2:
                    return text
        return ''

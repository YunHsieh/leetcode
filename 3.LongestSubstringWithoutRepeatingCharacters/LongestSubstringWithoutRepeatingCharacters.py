class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        start = -1; max_lenght = 0
        used_string = {}
        for index, char in enumerate(s):
            if used_string.get(char, -1) > start:
                start = used_string.get(char, 0)
            used_string[char] = index
            max_lenght = max(max_lenght, index-start)
        return max_lenght

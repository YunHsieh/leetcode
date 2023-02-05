from collections import Counter


class Solution:
    def findAnagrams(self, s: str, p: str) -> list[int]:
        p_counter, w = Counter(p), len(p)
        comp = Counter()
        result = []

        for i, text in enumerate(s):
            if text in p_counter:
                comp[text] += 1
            if i >= w and s[i-w] in p_counter:
                comp[s[i-w]] -= 1
            if comp == p_counter:
                result.append(i-w+1)
        return result

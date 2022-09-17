class Solution:
    def palindromePairs(self, words: List[str]) -> List[List[int]]:
        data = {v: i for i, v in enumerate(words)}
        results = []
        for i, word in enumerate(words):
            if (t := data.get(word[::-1])) not in [None, i]:
                results.append([i, t])
            if word == word[::-1] and (e := data.get('')) not in [None, i]:
                results.extend([[i, e], [e, i]])

            for j in range(1, len(word)):
                s1, s2 = word[j:], word[:j]
                if s1 == s1[::-1] and (index := data.get(s2[::-1])) is not None and index != i:
                    results.append([i, index])
                if s2 == s2[::-1] and (index := data.get(s1[::-1])) is not None and index != i:
                    results.append([index, i])
        return results

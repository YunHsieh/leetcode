class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        text_counter, w = Counter(s1), len(s1)   

        for i in range(len(s2)):
            if text_counter.get(s2[i]) is not None: 
                text_counter[s2[i]] -= 1
            # count from start
            if i >= w and text_counter.get(s2[i-w]) is not None: 
                text_counter[s2[i-w]] += 1

            if all([text_counter[i] == 0 for i in text_counter]):
                return True

        return False

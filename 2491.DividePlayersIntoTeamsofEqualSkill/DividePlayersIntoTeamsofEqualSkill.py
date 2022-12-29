class Solution:
    def dividePlayers(self, skill: list[int]) -> int:
        average = sum(skill) / (len(skill) / 2)
        result = 0
        mapping = {}
        for s in skill:
            mapping[s] = mapping.get(s, 0) + 1
        for k in mapping:
            if mapping[k] == 0:
                continue
            while mapping[k]:
                mapping[k] -= 1
                another = average - k
                if (tmp := mapping.get(another, None)):
                    mapping[another] -= 1
                if tmp is None or tmp <= 0:
                        return -1
                result += (another) * k
        return int(result)

class Solution:
    def smallestEquivalentString(self, s1: str, s2: str, baseStr: str) -> str:
        lexicograph = {}
        result = ''
        def union_find(data):
            if data == lexicograph.get(data, None):
                return data
            if not lexicograph.get(data, None) is None:
                return union_find(lexicograph.get(data))
            return data
        for a, b in zip(s1, s2):
            a, b = union_find(a), union_find(b)
            if a < b:
                a, b = b, a
            lexicograph[a] = b
        for i in baseStr:
            result += union_find(i)

        return result

class Solution:
    def findOriginalArray(self, changed: List[int]) -> List[int]:
        n = len(changed)
        if n % 2 != 0:
            return []
        changed = sorted(changed)
        data = {}
        for num in changed:
            data[num] = data.get(num, 0) + 1 
        results = []
        for k, v in data.items():
            if not v:
                continue
            if not data.get(k*2):
                return []
            if k == 0:
                results.extend([k]*(int(v/2)))
                continue
            results.extend([k]*v)
            data[k*2] -= v
            if data[k*2] < 0:
                return []
        return results

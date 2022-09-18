class Solution:
    def trap(self, height: List[int]) -> int:
        results = [None]*len(height)
        l, r, lh, rh = 0, len(height)-1, 0, 0
        
        while l < len(height):
            lh, rh = max(height[l], lh), max(height[r], rh)
            if results[l] is None:
                results[l], results[r] = lh - height[l], rh - height[r]
            else:
                results[l], results[r] = min(results[l], lh-height[l]), min(results[r], rh-height[r])
            if r==l:
                results[l] = min(min(results[l], lh-height[l]), min(results[r], rh-height[r]))
            l+=1
            r-=1
        return sum(results)

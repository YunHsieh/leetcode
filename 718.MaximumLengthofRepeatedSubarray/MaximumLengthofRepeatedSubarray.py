class Solution:
    def findLength(self, nums1: List[int], nums2: List[int]) -> int:
        results = [[0]*len(nums1) for _ in range(len(nums2))]
        results[0][0] = 1 if nums1[0] == nums2[0] else 0
        max_num = 0
        for i in range(1, len(results[0])):
            if nums1[i]==nums2[0]:
                results[0][i] = 1
        for i in range(1, len(results)):
            if nums1[0]==nums2[i]:
                results[i][0] = 1
        
        for i in range(1, len(results)):
            for j in range(1, len(results[0])):
                if nums1[j] == nums2[i]:
                    results[i][j] = results[i-1][j-1]+1
                else:
                    results[i][j] = 0
                max_num = max(results[i][j], max_num)
        return max_num

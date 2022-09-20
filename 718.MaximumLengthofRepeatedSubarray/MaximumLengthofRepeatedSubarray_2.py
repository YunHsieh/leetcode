"""
Runtime: 1733ms, 92.19%
Memory Usage: 14.3MB, 83.23%
"""

class Solution:
    def count_data(self, nums):
        data = []
        count = 1
        for i in range(len(nums)-1):
            if nums[i] == nums[i+1]:
                count += 1
            else:
                data.append([nums[i], count])
                count = 1
        data.append([nums[-1], count])
        return data

    def findLength(self, nums1: List[int], nums2: List[int]) -> int:
        nums1 = self.count_data(nums1)
        nums2 = self.count_data(nums2)
        last_arr = [0]*len(nums1)
        tmp = [0]*len(nums1)
        max_num = 0
        for i, d1 in nums2:
            count = 1
            tmp[0] = min(d1, nums1[0][1]) if i == nums1[0][0] else 0
            for j, d2 in nums1[1:]:
                if i == j:
                    tmp[count] = last_arr[count-1] + min(d1, d2)
                    if d1!=d2:
                        max_num = max(tmp[count], max_num)
                        tmp[count] = min(d1, d2)
                else:
                    tmp[count] = 0
                count+=1
            last_arr = tmp.copy()
            max_num = max(max(tmp), max_num)
        return max_num

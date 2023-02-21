class Solution:
    def singleNonDuplicate(self, nums: list[int]) -> int:
        n = len(nums)
        result = []
        if n == 1: return n
        if nums[0] != nums[1]: return nums[0]
        if nums[n-1] != nums[n-2]: return nums[n-1]
        def binary_search(s, e):
            mid = (e-s)//2 + s
            if nums[mid] not in {nums[mid-1], nums[mid+1]}:
                result.append(nums[mid])
            if result or e - mid <= 1:
                return 
            binary_search(mid, e)
            binary_search(s, mid)
        binary_search(1, n-1)
        return result[0]

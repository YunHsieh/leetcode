class Solution:
    def searchInsert(self, nums: list[int], target: int) -> int:
        n = len(nums)
        s, e = 0, n
        while True:
            mid = ((e-s)//2) + s
            if e-s == 1:
                if nums[s] >= target:
                    return s
                else:
                    return e
            if nums[mid] == target:
                return mid
            elif nums[mid] > target:
                e = mid
            else:
                s = mid

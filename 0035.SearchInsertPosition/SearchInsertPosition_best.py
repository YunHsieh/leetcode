class Solution:
    def searchInsert(self, nums: list[int], target: int) -> int:
        return sorted(nums + [target]).index(target)

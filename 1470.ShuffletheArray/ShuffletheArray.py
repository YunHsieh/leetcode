class Solution:
    def shuffle(self, nums: List[int], n: int) -> List[int]:
        result = []
        n = len(nums)
        for i in range(n//2):
            result.extend([nums[i], nums[(n//2)+i]])
        return result

class Solution:
    def maxSubarraySumCircular(self, nums: list[int]) -> int:
        total = result = sum(nums)
        cur_max = cur_min = 0

        for data in nums[1:]:
            cur_max = max(cur_max + data, data)
            cur_min = min(cur_min + data, data)
            result = max(result, cur_max, total - cur_min)
        return result

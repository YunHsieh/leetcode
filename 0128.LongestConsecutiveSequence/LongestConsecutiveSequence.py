class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        nums = set(nums)
        result = 0
        def find_consecutive(num, delta, data=0):
            if num+delta in nums:
                nums.remove(num+delta)
                return find_consecutive(num+delta, delta, data+1)
            return data
        while nums:
            num = nums.pop()
            result = max(result, find_consecutive(num, 1)+find_consecutive(num, -1)+1)
        return result

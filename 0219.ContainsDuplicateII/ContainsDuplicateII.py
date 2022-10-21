class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        record = {}
        for i, num in enumerate(nums):
            if num in record and k >= abs(record.get(num, i) - i):
                return True
            record[num] = i
        return False

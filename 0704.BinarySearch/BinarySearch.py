class Solution:
    def search(self, nums: List[int], target: int) -> int:
        head, tail = -1, len(nums)
        middle = 0
        while tail - head > 1:
            middle = (tail+head)//2
            if nums[middle] == target:
                return middle
            elif nums[middle] > target:
                tail = middle
            else:
                head = middle
        return -1
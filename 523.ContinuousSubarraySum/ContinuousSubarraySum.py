class Solution:
    def checkSubarraySum(self, nums: List[int], k: int) -> bool:
        data = {0:-1}
        result = 0
        for i, v in enumerate(nums):
            result = (result + v) % k
            if data.get(result, None) is not None:
                if i-data[result] >= 2:
                    return True
            else:
                data[result] = i
        return False

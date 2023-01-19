class Solution:
    def subarraysDivByK(self, nums: List[int], k: int) -> int:
        delta = result = 0
        groups = {0: 1}
        for num in nums:
           delta = (delta + num) % k
           result += groups.get(delta, 0)
           groups[delta] = groups.get(delta, 0) + 1

        return result

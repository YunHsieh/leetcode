class Solution:
    def sumEvenAfterQueries(self, nums: List[int], queries: List[List[int]]) -> List[int]:
        even = 0
        results = []
        for i in nums:
            if i % 2 == 0:
                even += i

        for v, i in queries:
            if nums[i] % 2 == 0:
                even -= nums[i]
            nums[i] += v
            if nums[i] % 2 == 0:
                even += nums[i]
            results.append(even)
        
        return results

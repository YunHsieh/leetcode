class Solution:
    def answerQueries(self, nums: list[int], queries: list[int]) -> list[int]:
        nums, n = sorted(nums), len(nums)
        results = []
        for i in range(1, n):
            nums[i] += nums[i-1]
        def binary_search(target, min_len, max_len):
            mid = int((max_len-min_len)/2) + min_len
            if nums[mid] > target and nums[mid-1] <= target:
                return mid
            elif nums[mid] > target:
                return binary_search(target, min_len, mid)
            else:
                return binary_search(target, mid, max_len)

        for v in queries:
            if v >= nums[-1]:
                results.append(n)
            elif v < nums[0]:
                results.append(0)
            else:
                results.append(binary_search(v, 0, n))
        
        return results

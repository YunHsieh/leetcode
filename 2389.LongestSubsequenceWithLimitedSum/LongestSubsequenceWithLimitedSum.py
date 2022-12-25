class Solution:
    def answerQueries(self, nums: list[int], queries: list[int]) -> list[int]:
        nums = sorted(nums)
        results = [0 for _ in range(len(queries))]
        for num in nums:
            for i, _ in enumerate(queries):
                if num <= queries[i]:
                    queries[i] -= num
                    results[i] += 1
        return results

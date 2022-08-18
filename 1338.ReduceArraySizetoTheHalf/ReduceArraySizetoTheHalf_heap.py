class Solution:
    def maxHeapify(self, arr, n, i):
        largest = i
        l = 2 * i + 1
        r = l + 1

        if l < n and self.number_count[arr[i]] < self.number_count[arr[l]]:
            largest = l

        if r < n and self.number_count[arr[largest]] < self.number_count[arr[r]]:
            largest = r

        if largest != i:
            arr[i], arr[largest] = arr[largest], arr[i]
            self.maxHeapify(arr, n, largest)

    def minSetSize(self, arr: List[int]) -> int:
        self.number_count = {}
        for i in arr:
            if self.number_count.get(i):
                self.number_count[i] += 1
            else:
                self.number_count[i] = 1
        diff_nums = list(self.number_count.keys())
        n = len(diff_nums)
        for i in range(n//2-1, -1, -1):
            self.maxHeapify(diff_nums, n, i)
        
        new_arr_count = 0
        for index, i in enumerate(range(n-1, 0, -1)):
            diff_nums[-1], diff_nums[0] = diff_nums[0], diff_nums[-1]
            new_arr_count += self.number_count[diff_nums.pop()]
            if new_arr_count >= len(arr) // 2:
                return index+1
            self.maxHeapify(diff_nums, i, 0)
        return 1

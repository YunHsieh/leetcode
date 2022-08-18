class Solution:

    def minSetSize(self, arr: List[int]) -> int:
        number_count = dict()
        for i in arr:
            number_count[i] = number_count.get(i, 0) + 1

        values = sorted([v for k, v in number_count.items()], reverse=True)
        result = 0
        for index, data in enumerate(values):
            result += data
            if result >= len(arr)//2:
                return index + 1
        return 1

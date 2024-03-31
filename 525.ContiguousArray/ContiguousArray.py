from typing import List


class Solution:
    def findMaxLength(self, nums: List[int]) -> int:
        existed_num = {0:-1}
        result = tmp = 0
        tmp = 0
        for i, num in enumerate(nums):
            if num == 0:
                tmp += 1
            else:
                tmp -= 1
            if (current:=existed_num.get(tmp)) != None:
                result = max(result, i - current)
            else:
                existed_num[tmp] = i
        return result

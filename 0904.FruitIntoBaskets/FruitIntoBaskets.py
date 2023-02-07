'''
Time complexity: O(N)
Space complexity: O(N)
Runtime: 84%+ 
Memory: 80%+
'''

class Solution:
    def totalFruit(self, fruits: list[int]) -> int:
        result = 1
        fruits_counter = {fruits[0]: 1}
        l, r = None, fruits[0]
        count = 1
        for i in fruits[1:]:
            if i != r:
                result = max(sum(fruits_counter.values()), result)
                if i not in {l, r}:
                    fruits_counter.pop(l, 0)
                    fruits_counter[r] = count
                count = 0
                l, r = r, i
            fruits_counter[i] = fruits_counter.get(i, 0) + 1
            count += 1
        return max(sum(fruits_counter.values()), result)

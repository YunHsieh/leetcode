class Solution:
    def dailyTemperatures(self, temperatures: list[int]) -> list[int]:
        res = [0] * len(temperatures)
        remain = []
        for i in range(len(temperatures)):
            while remain and temperatures[i] > temperatures[remain[-1]]:
                cur = remain.pop()
                res[cur] = i - cur
            remain.append(i)                    
        return res

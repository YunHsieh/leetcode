from typing import List


class Solution:
    def maximumUnits(self, boxTypes: List[List[int]], truckSize: int) -> int:
        ans = 0
        boxTypes = sorted(boxTypes, key=lambda item: item[1], reverse=True)
        print(boxTypes)
        for box, item in boxTypes:
            if truckSize < box:
                ans += truckSize * item
            else:
                ans += box * item
            truckSize -= box
            if truckSize < 0:
                break
            
        return ans

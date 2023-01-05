class Solution:
    def findMinArrowShots(self, points: list[list[int]]) -> int:
        points.sort()
        cur_max = points[0][1]
        result = 1
        for point in points[1:]:
            if cur_max >= point[0]:
                cur_max = min(cur_max, point[1])
            else:
                cur_max = point[1]
                result+=1
        return result

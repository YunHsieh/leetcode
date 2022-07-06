"""
. = smooth fragments
x = potholes
When you fix a pothole and that can make to K+1 costs.

solution(s, budget)
For instance:
s = ...xxx..x....xxx
budget = 7
return 5
explaintation:
fixed to ........x....x..
So fixed the totally five potholes.

"""

import re
def solution(s, b):
    result = 0
    potholes = re.findall(r'x+', s)
    potholes = sorted(potholes, reverse=True)

    for i in potholes:
        n = len(i) + 1
        if b-n > 0:
            result += n-1
            b -= n
        else:
            result += b-1
            break
    return result


assert solution('...xxx..x', 5) == 3
assert solution('...xxx..x....xxx', 7) == 5
assert solution('..xxxxx', 4) == 3
assert solution('x.x.xxx...x', 14) == 6
assert solution('..', 5) == 0

"""
Given 10 positions. A B C  D E F G  H I J K
You need to find 4 seats in a row and count it.
There are 2 aisles in between C and D, G and H.
When you allocate the seat you can't let one for single.

def solutions(n: int, limit: str) -> int

For instances:
n=2 limit="1A 2F 1C" , return = 2
explaintation:
Since 1A 2F 1C can not be aloocated.
So [1D 2E 1F 1G] [1B 1C 1D 1E] = 2

n=1 limit="" , return = 2

"""


def has_seat(seat_a, seat_b, prohibit):
    for i in range(4):
        if f"{seat_a}{chr(seat_b + i + ord('A'))}" in prohibit:
            return False
    return True
    

def solution(n: int, s: str) -> int:
    result = 0
    s_list = s.split(' ')
    start_seat = [1, 3, 5]
    for i in range(n):
        j = 0
        while j < 3:
            if has_seat(i+1, start_seat[j], s_list):
                j+=1
                result += 1
            j+=1
    return result

assert solution(2, '1A 2F 1C') == 2
assert solution(1, '') == 2

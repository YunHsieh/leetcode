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


def solution(n: int, s: str) -> int:
    result = 0
    s_list = s.split(' ')
    seat = [[0]*11 for i in range(n)]
    for i in s_list:
        right_number, left_number = int(i[0]), ord(i[1] - ord('A'))
        seat[right_number][left_number] = 1

    for i, data in enumerate(seat):
        for j in [0, 3, 5]:
            count = 0
            verify = True
            while count < 4:
                if data[j+count] == 1:
                    verify=False
                    break
                data[j+count] = 1
            if verify:
                while count < 4:
                    data[j+count] == 1
    return result

assert solution(2, '1A 2F 1C') == 2
assert solution(1, '') == 2

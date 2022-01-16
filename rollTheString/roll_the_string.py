"""
if give the 'abz' and [3,2,1]
abz 3
-> b c a
bca 2
-> c d a
cda 1
-> dda
"""

def roll_the_string(s, roll):
    if not s or not roll:
        return s

    s = list(s)
    for _index, _val in enumerate(roll):
        for _i in range(_val):
            if ord(s[_i%len(s)])+1 > 122:
                s[_i%len(s)] = chr(97)
            else:
                s[_i%len(s)] = chr(ord(s[_i%len(s)])+1)
    return ''.join(s)

def solution(A):
    ans = 0
    start = 0
    while start < len(A):
        if (A[start-1] + A[start]) % 2 == 0:
            ans += 1
            start += 1
        start += 1
    return ans

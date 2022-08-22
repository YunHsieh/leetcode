class Solution:
    def isPowerOfFour(self, n: int) -> bool:
        if n == 1:
            return True
        elif n == 0:
            return False

        if (n := n/4) == int(n):
            return self.isPowerOfFour(n)
        else:
            return False

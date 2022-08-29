from typing import List

class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        self.grid = grid
        self.m, self.n = len(grid), len(grid[0])
        return sum(self.eraseIslands(i, j) for i in range(self.m) for j in range(self.n))

    def eraseIslands(self, x, y):
        if (x<0 
            or y<0 
            or x==self.m 
            or y==self.n 
            or self.grid[x][y]=='0'):
            return 0

        if self.grid[x][y]=='1':
            self.grid[x][y] = '0'
            self.eraseIslands(x+1, y)
            self.eraseIslands(x-1, y)
            self.eraseIslands(x, y+1)
            self.eraseIslands(x, y-1)
            return 1

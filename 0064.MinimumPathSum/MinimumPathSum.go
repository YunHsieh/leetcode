package main

func MinPathSum(grid [][]int) int {
    n := len(grid)
    nj := len(grid[0])
    for i:=1; i<len(grid[0]); i++ {
        grid[0][i] += grid[0][i-1]
    }
    for i:=1; i<len(grid); i++ {
        grid[i][0] += grid[i-1][0]
    }
    for i:=1; i<len(grid); i++ {
        for j:=1; j<len(grid[0]); j++ {
            if grid[i][j-1] > grid[i-1][j] {
                grid[i][j] += grid[i-1][j]
            } else {
                grid[i][j] += grid[i][j-1]
            }
        }
    }
    return grid[n-1][nj-1]
}

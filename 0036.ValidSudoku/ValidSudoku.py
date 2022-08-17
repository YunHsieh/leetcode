from typing import List


class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        lines = [{} for i in range(9)]
        cols = [{} for i in range(9)]
        grids = [{} for i in range(9)]
        for i, data in enumerate(board):
            for j, num in enumerate(data):
                if num == ".":
                    continue
                grids_index = (i//3*3)+(j//3)
                if any([
                    lines[i].get(num),
                    cols[j].get(num),
                    grids[grids_index].get(num),
                ]):
                    return False
                lines[i][num] = 1
                cols[j][num] = 1
                grids[grids_index][num] = 1
        return True

class Solution:
    def canVisitAllRooms(self, rooms: list[list[int]]) -> bool:
        if (n := len(rooms)) < 1:
            return True
        keys = set(rooms[0])
        visited = {0}

        def visit_room():
            enter = keys - visited
            if not enter:
                return False
            next_room = list(enter)[0]
            keys.update(rooms[next_room])
            visited.add(next_room)
            if len(visited) == n:
                return True
            return visit_room()
        return visit_room()

input = """RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE"""

with open('input.txt', 'r') as file:
    input = file.read()

map = [[i for i in line] for line in input.strip().split("\n")]

n = len(map)
m = len(map[0])

visited = [[False for _ in range(m)] for _ in range(n)]

def get(x, y):
    if 0 <= x < m and 0 <= y < n:
        return map[y][x]

def get2(x, y):
    if 0 <= x < m and 0 <= y < n:
        return visited[y][x]
    return True

def parcours(x, y):
    if get(x, y) is None:
        return 0, 0
    visited[y][x] = True
    p, a = 4, 1
    for dx, dy in [(1, 0), (0, 1), (-1, 0), (0, -1)]:
        if get(x+dx, y+dy) == get(x, y):
            p -= 1
            if not get2(x+dx, y+dy):
                result = parcours(x+dx, y+dy)
                p += result[0]
                a += result[1]
    return p, a

score = 0
for y in range(n):
    for x in range(m):
        if not get2(x, y):
            result = parcours(x, y)
            score += result[0] * result[1]

print(score)

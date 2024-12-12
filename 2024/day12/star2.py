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

# input = """AAAAAA
# AAABBA
# AAABBA
# ABBAAA
# ABBAAA
# AAAAAA"""

with open('input.txt', 'r') as file:
    input = file.read()

map = [[i for i in line] for line in input.strip().split("\n")]

n = len(map)
m = len(map[0])

visited = [[False for _ in range(m)] for _ in range(n)]

def get_map(x, y):
    if 0 <= x < m and 0 <= y < n:
        return map[y][x]

def get_visited(x, y):
    if 0 <= x < m and 0 <= y < n:
        return visited[y][x]
    return True

def ortho(dx, dy):
    if dx != 0:
        return [((0, 1), (dx, 1)), ((0, -1), (dx, -1))]
    else:
        return [((1, 0), (1, dy)), ((-1, 0), (-1, dy))]

def aff(x, y):
    print('\n'.join([''.join([c if (x!=i or y!=j) and get_visited(i, j) else 'X' if x==i and y==j else ' ' for i, c in enumerate(line)]) for j, line in enumerate(map)]))
    print('----')

def parcours(x, y):
    # aff(x, y)
    if get_map(x, y) is None:
        return 0, 0
    visited[y][x] = True
    p, a = 0, 1
    # print(f"{x}, {y}")
    # aff(x, y)
    for dx, dy in [(1, 0), (0, 1), (-1, 0), (0, -1)]:
        if get_map(x+dx, y+dy) != get_map(x, y):
            dp = 1
            for (dx1, dy1), (dx2, dy2) in ortho(dx, dy):
                if get_map(x+dx1, y+dy1) == get_map(x, y) and get_visited(x+dx1, y+dy1) and get_map(x+dx2, y+dy2) != get_map(x, y):
                    dp -= 1
            p += dp
            if dp < 0:
                print(f"decreased a lot {x}, {y}")
    for dx, dy in [(1, 0), (0, 1), (-1, 0), (0, -1)]:
        if get_map(x+dx, y+dy) == get_map(x, y):
            if not get_visited(x+dx, y+dy):
                result = parcours(x+dx, y+dy)
                p += result[0]
                a += result[1]
    # print(p, a)
    return p, a

score = 0
for y in range(n):
    for x in range(m):
        if not get_visited(x, y):
            result = parcours(x, y)
            print(x, y, result)
            score += result[0] * result[1]

print(score)

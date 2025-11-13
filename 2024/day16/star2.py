import queue

input = """###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############"""
# input = """#################
# #...#...#...#..E#
# #.#.#.#.#.#.#.#.#
# #.#.#.#...#...#.#
# #.#.#.#.###.#.#.#
# #...#.#.#.....#.#
# #.#.#.#.#.#####.#
# #.#...#.#.#.....#
# #.#.#####.#.###.#
# #.#.#.......#...#
# #.#.###.#####.###
# #.#.#...#.....#.#
# #.#.#.#####.###.#
# #.#.#.........#.#
# #.#.#.#########.#
# #S#.............#
# #################"""

# with open('input.txt', 'r') as file:
#     input = file.read()

raw_map = input.strip().split("\n")

w, h = len(raw_map[0]), len(raw_map)
map = []
starting_pos = (0, 0)
end_pos = (0, 0)
for y, raw_row in enumerate(raw_map):
    row = []
    for x, c in enumerate(raw_row):
        if c == "S":
            starting_pos = (x, y)
            row.append(0)
        elif c == "E":
            end_pos = (x, y)
            row.append(0)
        elif c == ".":
            row.append(0)
        else:
            row.append(1)
    map.append(row)

x, y = starting_pos

def get_char(cx, cy, path=None):
    if path and (cx, cy) in path:
        return "X"
    elif (cx, cy) == starting_pos:
        return "S"
    elif (cx, cy) == end_pos:
        return "E"
    elif map[cy][cx] == 0:
        return " "
    else:
        return "#"

def print_map(path):
    print('\n'.join([''.join([get_char(mx, my, path) for mx in range(w)]) + " " + ''.join([get_char(mx, my) for mx in range(w)]) for my in range(h)]))
    print('-'*len(map[0]))

# dirs :
#   1
# 4   2
#   3

def next_tile(x, y, dir):
    nx, ny = x + (1 if dir == 2 else -1 if dir == 4 else 0), y + (1 if dir == 3 else -1 if dir == 1 else 0)
    if 0 <= nx < w and 0 <= ny < h:
        return nx, ny

best_path = None
paths = set()

q = queue.PriorityQueue()
v = [[[(None, ())] * 5 for _ in range(w)] for _ in range(h)]

while not q.empty():
    p, (x, y, dir, c) = q.get()
    # print(x, y, dir)
    if (x, y) == end_pos:
        print("found best path", p)
        if best_path is None or best_path > p:
            best_path = p
            paths = set(c)
            print_map(c)
        elif best_path == p:
            paths.update(c)
            print_map(c)
        # else:
        #     break

    if not v[y][x][dir]:
        v[y][x][dir] = True
        nt = next_tile(x, y, dir)
        if nt and map[nt[1]][nt[0]] == 0:
            nx, ny = nt
            nc = c.copy()
            nc.append((nx, ny))
            q.put((p+1, (nx, ny, dir, nc)))
        for d in range(1, 5):
            if not v[y][x][d]:
                q.put((p+1000, (x, y, d, c)))
        # v[y][x][dir] = True
        # for d in range(1, 5):
        #     nt = next_tile(x, y, d)
        #     # print(c)
        #     if nt and map[nt[1]][nt[0]] == 0:
        #         nx, ny = nt
        #         nc = c.copy()
        #         nc.add((nx, ny))
        #         q.put((p+(1 if d==dir else 1001), (nx, ny, d, nc)))

print(len(paths))

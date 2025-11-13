input = """....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#..."""

with open('input.txt', 'r') as file:
    input = file.read()

map = []
starting_pos = None
for i, line in enumerate(input.strip().split("\n")):
    map.append([c for c in line])
    if "^" in line:
        starting_pos = (i, line.index("^"))
        # map[starting_pos[0]][starting_pos[1]] = "X"
print(starting_pos)
print(map)

n, m = len(map), len(map[0])

orientation = (-1, 0)

def rotate(orientation):
    y, x = orientation
    return (x, -y)

pos = starting_pos

trajectory = set()

while 0 <= pos[0] < n and 0 <= pos[1] < m:
    trajectory.add(pos)
    # map[pos[0]][pos[1]] = "X"
    next_pos = (pos[0]+orientation[0], pos[1]+orientation[1])
    # print(next_pos, orientation)
    if 0 <= next_pos[0] < n and 0 <= next_pos[1] < m and map[next_pos[0]][next_pos[1]] == "#":
        orientation = rotate(orientation)
    else:
        pos = next_pos

trajectory.remove(starting_pos)

def deep_copy(map):
    return [[i for i in row] for row in map]

def circling(map):
    visited = {}
    pos = starting_pos
    orientation = (-1, 0)
    while 0 <= pos[0] < n and 0 <= pos[1] < m:
        if (pos, orientation) in visited:
            return True
        visited[(pos, orientation)] = True
        next_pos = (pos[0]+orientation[0], pos[1]+orientation[1])
        if 0 <= next_pos[0] < n and 0 <= next_pos[1] < m and map[next_pos[0]][next_pos[1]] == "#":
            orientation = rotate(orientation)
        else:
            pos = next_pos
    return False

count = 0
total = len(trajectory)
current = 0
for block in trajectory:
    current += 1
    if current % 100 == 0:
        print(f"{current} / {total}")
    current_map = deep_copy(map)
    current_map[block[0]][block[1]] = "#"
    if circling(current_map):
        count += 1

print(count)

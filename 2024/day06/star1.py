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

while 0 <= pos[0] < n and 0 <= pos[1] < m:
    # print(pos, orientation)
    map[pos[0]][pos[1]] = "X"
    next_pos = (pos[0]+orientation[0], pos[1]+orientation[1])
    print(next_pos, orientation)
    if 0 <= next_pos[0] < n and 0 <= next_pos[1] < m and map[next_pos[0]][next_pos[1]] == "#":
        orientation = rotate(orientation)
    else:
        pos = next_pos

print(sum([row.count("X") for row in map]))

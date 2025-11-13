input = """##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^"""

with open('input.txt', 'r') as file:
    input = file.read()

raw_map, moves = input.strip().split("\n\n")
raw_map_rows = raw_map.split("\n")

w, h = len(raw_map_rows[0]), len(raw_map_rows)
map = []
starting_pos = (0, 0)
for y, raw_row in enumerate(raw_map_rows):
    row = []
    for x, c in enumerate(raw_row):
        if c == "@":
            starting_pos = (x, y)
            row.append(0)
        elif c == ".":
            row.append(0)
        elif c == "O":
            row.append(1)
        else:
            row.append(-1)
    map.append(row)

x, y = starting_pos
def print_map():
    print('\n'.join([''.join(['@' if x == mx and y == my else '#' if c == -1 else 'O' if c == 1 else '.' for mx, c in enumerate(row)]) for my, row in enumerate(map)]))
    print('-'*len(map[0]))

print(starting_pos)
for m in moves:
    # print(m)
    # print_map()
    if m == "\n":
        continue
    if m == "<":
        dir = (-1, 0)
    elif m == "^":
        dir = (0, -1)
    elif m == ">":
        dir = (1, 0)
    else:
        dir = (0, 1)

    dx, dy = dir
    px, py = x+dx, y+dy
    while map[py][px] != 0:
        if map[py][px] == -1: # wall
            # print('broke', px, py)
            break
        px, py = px + dx, py + dy
    else:
        # print(px, py)
        map[py][px] = 1
        map[y+dy][x+dx] = 0
        x += dx
        y += dy

print_map()

gps = 0
for y in range(1, len(map)):
    for x in range(1, len(map[0])):
        if map[y][x] == 1:
            gps += x + y * 100

print(x, y)
print(gps)

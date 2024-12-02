import math

import sys
sys.setrecursionlimit(10000000)

with open("input_check.txt", 'r', encoding='utf-8') as file:
    data = [line.strip() for line in file.readlines()]

PIPES = {
    "|": ("N", "S"),
    "-": ("E", "W"),
    "L": ("N", "E"),
    "J": ("N", "W"),
    "7": ("S", "W"),
    "F": ("S", "E"),
    "S": ("E", "S"),
}
DIRECTIONS = {
    "N": (0, -1),
    "S": (0, 1),
    "W": (-1, 0),
    "E": (1, 0),
}

class Grid:
    def __init__(self, data: list, main: bool = False):
        self.data = data
        self.max_x = len(data[0])
        self.max_y = len(data)
        if main:
            self.distances = Grid([[math.inf for _ in range(self.max_x)] for _ in range(self.max_y)])
        self.main = main
        print(f"The grid is of size {self.max_x}, {self.max_y}")

    def __getitem__(self, pos):
        x, y = pos
        if (0 <= x < self.max_x) and (0 <= y < self.max_y):
            return self.data[y][x]
        else:
            if self.main:
                return "."
            else:
                return -1

    def __setitem__(self, pos, value):
        x, y = pos
        if (0 <= x < self.max_x) and (0 <= y < self.max_y):
            self.data[y][x] = value

    def __iter__(self):
        for y in range(self.max_y):
            for x in range(self.max_x):
                yield x, y

    def directions(self, pos):
        """Yield all positions to go to"""
        x, y = pos
        if self[x, y] in PIPES:
            for direction in PIPES[self[x, y]]:
                d_x, d_y = DIRECTIONS[direction]
                target_x, target_y = x + d_x, y + d_y
                yield target_x, target_y

    def __repr__(self):
        if self.main:
            return '\n'.join(
                [''.join([str(i) if i != math.inf else " " for i in self.data[y]]) for y in range(self.max_y)]
                )
        else:
            return '\n'.join(
                [''.join([str(i) if i != math.inf else " " for i in self.data[y]]) for y in range(self.max_y)]
                )

grid = Grid(data, main=True)

start_x, start_y = 0, 0
for x, y in grid:
    if grid[x, y] == "S":
        start_x, start_y = x, y
        break

def deep(pos, length: int, grid: Grid):
    x, y = pos
    for next_x, next_y in grid.directions((x, y)):
        if grid.distances[next_x, next_y] > length + 1:
            grid.distances[next_x, next_y] = length + 1
            deep((next_x, next_y), length + 1, grid)

def print_loop(grid: Grid):
    output = ""
    for y in range(grid.max_y):
        for x in range(grid.max_x):
            if grid.distances[x, y] != math.inf:
                output += grid[x, y]
            else:
                output += " "
        output += "\n"
    return output

print("Start:", start_x, start_y)

# We need to figure which pipe the start connects to
# and because I don't want to code it, I'm going to hardcode it
#DIRECTIONS["S"] = ("N", "S")
#DIRECTIONS["S"] = ("E", "S")

print(grid)
print("")
grid.distances[start_x, start_y] = 0
deep((start_x, start_y), 0, grid)
print(print_loop(grid))
with open("output.txt", 'w', encoding='utf-8') as file:
    file.write(print_loop(grid))

max = 0
max_pos = start_x, start_y
for x, y in grid.distances:
    distance = grid.distances[x, y]
    if distance != math.inf and distance > max:
        max = grid.distances[x, y]
        max_pos = (x, y)

print(f"Max: {max}")
print(f"Max pos: {max_pos}")

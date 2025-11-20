with open("input.txt", 'r', encoding='utf-8') as file:
    data = [line.strip() for line in file.readlines()]

class Grid:
    def __init__(self, data: list):
        self.data = data
        self.max_x = len(data[0])
        self.max_y = len(data)
        print(f"The grid is of size {self.max_x}, {self.max_y}")

    def __getitem__(self, pos):
        x, y = pos
        if (0 <= x < self.max_x) and (0 <= y < self.max_y):
            return self.data[y][x]
        else:
            return "."

    def __setitem__(self, pos, value):
        x, y = pos
        if (0 <= x < self.max_x) and (0 <= y < self.max_y):
            self.data[y] = self.data[y][:x] + value + self.data[y][x+1:]
        else:
            print("out of grid")

    def __iter__(self):
        for y in range(self.max_y):
            for x in range(self.max_x):
                yield x, y

    def first_rock(self, pos):
        x, y = pos
        while self[x, y-1] == "." and 0 < y < self.max_y:
            y -= 1
        return x, y

    def __repr__(self):
        return '\n'.join(
            [self.data[y] for y in range(self.max_y)]
        )

grid = Grid(data)

for y in range(grid.max_y):
    for x in range(grid.max_x):
        if grid[x, y] == "O":
            x_target, y_target = grid.first_rock((x, y))
            grid[x, y] = "."
            grid[x_target, y_target] = "O"

print(grid)

sum = 0
for y in range(grid.max_y):
    for x in range(grid.max_x):
        if grid[x, y] == "O":
            sum += grid.max_y - y

print(sum)


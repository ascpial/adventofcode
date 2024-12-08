input = """............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............"""

with open('input.txt', 'r') as file:
    input = file.read()

data = {}

lines = input.strip().split("\n")
height = len(lines)
width = len(lines[0])

for y, line in enumerate(lines):
    for x, char in enumerate(line):
        if char != ".":
            if char in data:
                data[char].add((y, x))
            else:
                data[char] = {(y, x)}

targetted = [[False for i in range(width)] for j in range(height)]

def set(y, x):
    if 0 <= y < height and 0 <= x < width:
        targetted[y][x] = True

for frequency, antenas in data.items():
    for y1, x1 in antenas:
        for y2, x2 in antenas:
            if (y1, x1) != (y2, x2):
                dy = abs(y2-y1)
                dx = abs(x2-x1)

                x, y = x1, y1
                while 0 <= x < width and 0 <= y < height:
                    set(y, x)
                    y = y + (dy if y1 > y2 else -dy)
                    x = x + (dx if x1 > x2 else -dx)

                x, y = x2, y2
                while 0 <= x < width and 0 <= y < height:
                    set(y, x)
                    y = y + (dy if y1 < y2 else -dy)
                    x = x + (dx if x1 < x2 else -dx)

print(sum([sum([1 if ant else 0 for ant in line]) for line in targetted]))
